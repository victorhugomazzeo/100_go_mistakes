// Mistake #28: Maps and memory leaks
//
// A map never shrinks: deleting every key keeps the space it grew to. To give
// that memory back, re-create the map (make a new one, copy what you keep).
package main

import (
	"fmt"
	"runtime"
)

const n = 1_000_000

func main() {
	mistake()
	fix()
}

func mistake() {
	m := make(map[int][128]byte)
	for i := range n {
		m[i] = [128]byte{}
	}
	runtime.GC() // clears the old tables left over from growing, so we measure only the map
	printAlloc("mistake: after adding")

	for i := range n {
		delete(m, i) // empties the slot, but the map keeps its size
	}
	runtime.GC()
	printAlloc("mistake: after deleting") // memory is still held by the empty map
	runtime.KeepAlive(m)
}

func fix() {
	m := make(map[int][128]byte)
	for i := range n {
		m[i] = [128]byte{}
	}
	runtime.GC() // same as in mistake(): measure only the map
	printAlloc("fix: after adding")

	for i := range n {
		delete(m, i)
	}
	m = make(map[int][128]byte) // a new, small map: the old big one has no references left
	runtime.GC()
	printAlloc("fix: after recreating")
	runtime.KeepAlive(m)
}

func printAlloc(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %d MB\n", label, m.Alloc/1024/1024)
}
