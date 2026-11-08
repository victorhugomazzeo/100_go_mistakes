// Mistake #24: Not making slice copies correctly
//
// copy(dst, src) copies min(len(dst), len(src)) elements: copying into an
// empty dst copies nothing. Make dst with len(src) first, or use slices.Clone.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	src := []int{0, 1, 2}
	var dst []int
	n := copy(dst, src)
	fmt.Printf("mistake: n:%d dest:%v\n", n, dst)
}

func fix() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("fix: n:%d dest:%v\n", n, dst)
}
