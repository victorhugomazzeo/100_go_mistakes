// Mistake #33: Making wrong assumptions during map iterations
//
// Map iteration order is random, and keys added during a range may or may
// not be visited. Sort the keys for a stable order; write into a copy.
package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	// 1. Order: range visits the keys in a random order, run after run.
	letters := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Print("mistake: order: ")
	for k := range letters {
		fmt.Print(k, " ")
	}
	fmt.Println()

	// 2. Insert while iterating: a key added during the range may or may not
	// be visited by this same loop. Expected 5 keys (0, 1, 2, 10, 12), but if
	// the loop visits 10, it adds 20, which may add 30... different every run.
	m := map[int]bool{0: true, 1: false, 2: true}
	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}
	fmt.Println("mistake: insert:", len(m), "keys:", m)
}

func fix() {
	// 1. Order: sort the keys first, then range over the sorted slice.
	letters := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Print("fix: order: ")
	for _, k := range slices.Sorted(maps.Keys(letters)) {
		fmt.Print(k, " ")
	}
	fmt.Println()

	// 2. Insert while iterating: read from m, write into a copy. The loop only
	// sees m, which never changes, so it visits exactly the 3 original keys.
	m := map[int]bool{0: true, 1: false, 2: true}
	m2 := maps.Clone(m)
	for k, v := range m {
		if v {
			m2[10+k] = true
		}
	}
	fmt.Println("fix: insert:", len(m2), "keys:", m2)
}
