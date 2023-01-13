package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from the Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Defining a counter
	wc := 0

	switch {
	case countLines:
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			wc++
		}
	case countBytes:
		for {
			if _, err := r.Read(make([]byte, 1)); err == io.EOF {
				break
			}
			wc++
		}
	default:
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wc++
		}
	}
	return wc
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
