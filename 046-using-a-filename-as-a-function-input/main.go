// Mistake #46: Using a filename as a function input
//
// Taking a filename ties a function to files and forces tests to create them.
// Take an io.Reader: files, strings and HTTP bodies all fit, and tests are easy.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	n, err := countEmptyLinesInFile("notes.txt")
	fmt.Println("mistake:", n, err)
}

func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close() // the #35: closes when the function returns

	count := 0
	scanner := bufio.NewScanner(file) // reads line by line
	for scanner.Scan() {              // advances to the next line, false at the end
		if scanner.Text() == "" { // the current line, as a string
			count++
		}
	}
	return count, scanner.Err()
}

func countEmptyLines(r io.Reader) (int, error) {
	count := 0
	scanner := bufio.NewScanner(r) // reads from whatever r is: file, string, HTTP body...
	for scanner.Scan() {
		if scanner.Text() == "" {
			count++
		}
	}
	return count, scanner.Err()
}

func fix() {
	text := "first\n\nsecond\n\n\nthird"
	n, err := countEmptyLines(strings.NewReader(text))
	fmt.Println("fix:", n, err)
}
