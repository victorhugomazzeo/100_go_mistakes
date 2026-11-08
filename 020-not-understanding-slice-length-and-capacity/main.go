// Mistake #20: Not understanding slice length and capacity
//
// A slice points to a backing array: len is how many elements you can use,
// cap is how many the array holds. append fills cap, then grows a new array.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	s := make([]int, 3, 6)
	printSlice(s)
	// s[4] = 0 // panic: index out of range [4] with length 3: the index must be < len, not cap
}

func fix() {
	s := make([]int, 3, 6)
	printSlice(s) // 1. before anything

	s = append(s, 2)
	printSlice(s) // 2. fits in cap: same array

	s = append(s, 3, 4, 5)
	printSlice(s) // 3. exceeds cap: new array, cap doubled
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d ptr=%p %v\n", len(s), cap(s), s, s)

}
