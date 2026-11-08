// Mistake #21: Inefficient slice initialization
//
// When the final size is known, create the slice with that capacity (or
// length): make([]int, 0, n) avoids reallocating and copying as it grows.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	n := 10
	s := make([]int, 0)
	for i := range n {
		s = append(s, i)
		fmt.Printf("mistake: len=%d cap=%d ptr=%p\n", len(s), cap(s), s)
	}
}
func fix() {
	n := 10
	s := make([]int, 0, n)
	for i := range n {
		s = append(s, i)
		fmt.Printf("fix: len=%d cap=%d ptr=%p\n", len(s), cap(s), s)
	}
}
