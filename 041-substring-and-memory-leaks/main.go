// Mistake #41: Substring and memory leaks
//
// A substring (log[:36]) shares the bytes of the original string, so keeping
// it keeps the whole string alive. Use strings.Clone to keep only the piece.
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	id := receiveLog()[:36]
	runtime.GC()
	printAlloc("mistake")
	runtime.KeepAlive(id)
}

func fix() {
	id := strings.Clone(receiveLog()[:36])
	runtime.GC()
	printAlloc("fix")
	runtime.KeepAlive(id)
}

// receiveLog simulates a big log line: a 36-char ID followed by 10 MB.
func receiveLog() string {
	return "123e4567-e89b-12d3-a456-426614174000" + strings.Repeat("x", 10_000_000)
}

func printAlloc(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %d KB\n", label, m.Alloc/1024)
}
