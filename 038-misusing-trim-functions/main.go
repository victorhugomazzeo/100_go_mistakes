// Mistake #38: Misusing trim functions
//
// TrimRight removes trailing runes found in a set, not a suffix:
// TrimRight("logo.go", ".go") is "l". To drop a suffix, use TrimSuffix.
package main

import (
	"fmt"
	"strings"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	for _, file := range []string{"main.go", "logo.go"} {
		fmt.Println("mistake:", strings.TrimRight(file, ".go"))
	}
}

func fix() {
	for _, file := range []string{"main.go", "logo.go"} {
		fmt.Println("fix:", strings.TrimSuffix(file, ".go"))
	}
}
