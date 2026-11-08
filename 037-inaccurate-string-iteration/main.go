// Mistake #37: Inaccurate string iteration
//
// range over a string walks rune by rune, but i is a byte index: s[i] is one
// byte, so "é" prints as "Ã". Use the rune from range: for i, r := range s.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	s := "café"
	for i := range s {
		fmt.Printf("mistake: position %d: %c\n", i, s[i])
	}
}

func fix() {
	s := "café"
	for i, r := range s {
		fmt.Printf("fix: position %d: %c\n", i, r)
	}
}
