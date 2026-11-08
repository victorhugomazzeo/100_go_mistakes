// Mistake #27: Inefficient map initialization
//
// When the number of entries is known, give it to make: make(map[K]V, n)
// sizes the map once, instead of growing and moving entries as it fills.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	mistake()
	fix()
}

// mallocs returns how many memory allocations the program has made so far.
func mallocs() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Mallocs
}

const n = 1_000_000

func mistake() {
	before := mallocs()
	start := time.Now()

	m := make(map[int]int) // no size: the map grows as it fills
	for i := range n {
		m[i] = i
	}

	fmt.Println("mistake:", time.Since(start), mallocs()-before, "allocations")
}

func fix() {
	before := mallocs()
	start := time.Now()

	m := make(map[int]int, n) // sized up front: no growth while filling

	for i := range n {
		m[i] = i
	}

	fmt.Println("fix:", time.Since(start), mallocs()-before, "allocations")
}
