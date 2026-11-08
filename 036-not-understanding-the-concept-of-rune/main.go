// Mistake #36: Not understanding the concept of rune
//
// len(s) counts bytes, not characters: in UTF-8 "é" takes 2 bytes, so
// len("café") is 5. A rune is one character (code point); count runes instead.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	s := "café"
	fmt.Println("mistake: letters in", s, len(s))
	fmt.Printf("mistake: bytes of %s: % x\n", s, s)
}

func fix() {
	s := "café"
	n := utf8.RuneCountInString(s)
	fmt.Println("fix: letters in", s, n)
	fmt.Printf("fix: bytes of %s: % x\n", s, s)
}
