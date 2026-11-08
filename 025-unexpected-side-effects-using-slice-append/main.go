// Mistake #25: Unexpected side effects using slice append
//
// s[:2] shares s's array and keeps its spare cap, so appending to it
// overwrites s[2]. Cap it with s[:2:2] (or copy) so append allocates anew.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	s := []int{1, 2, 3}
	part := s[:2] // len 2, cap 3: part can still see the slot of s[2]

	part = append(part, 10) // fits in cap: writes 10 into s's array, over the 3
	fmt.Println("mistake: part:", part, "cap:", cap(part))
	fmt.Println("mistake: s:   ", s)
}

func fix() {
	s := []int{1, 2, 3}
	part := s[:2:2] // len 2, cap 2 (max - low): part can't see s[2]

	part = append(part, 10) // no room in cap: allocates a new array, s is untouched
	fmt.Println("fix: part:", part, "cap:", cap(part))
	fmt.Println("fix: s:   ", s)
}
