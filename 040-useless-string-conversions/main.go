// Mistake #40: Useless string conversions
//
// Converting []byte to string and back copies the data. The bytes package
// mirrors strings (TrimSpace, Split, Contains...), so stay in []byte.
package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func main() {
	mistake()
	fix()
}

// allocatedKB returns how many KB the program has allocated so far.
func allocatedKB() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.TotalAlloc / 1024
}

func mistake() {
	b := []byte("   " + strings.Repeat("x", 1_000_000) + "   ")

	before := allocatedKB()
	trimmed := []byte(strings.TrimSpace(string(b))) // []byte → string → []byte
	fmt.Println("mistake:", allocatedKB()-before, "KB copied, len", len(trimmed))
}

func fix() {
	b := []byte("   " + strings.Repeat("x", 1_000_000) + "   ")

	before := allocatedKB()
	trimmed := bytes.TrimSpace(b) // stays in []byte: no conversion
	fmt.Println("fix:", allocatedKB()-before, "KB copied, len", len(trimmed))
}
