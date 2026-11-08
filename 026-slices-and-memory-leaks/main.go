// Mistake #26: Slices and memory leaks
//
// A small slice cut from a big one (msg[:5]) keeps the whole backing array
// alive, so the GC can't free it. Copy the few bytes you need instead.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	mistake()
	fix()
}

// receiveMessage simulates a big message: 10 MB.
func receiveMessage() []byte {
	return make([]byte, 10_000_000)
}

// printAlloc prints how much memory is in use right now.
func printAlloc(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: %d KB\n", label, m.Alloc/1024)
}

func mistake() {
	msgType := receiveMessage()[:5] // only 5 bytes from a 10 MB message
	runtime.GC()                    // forces the GC to run now
	printAlloc("mistake")
	runtime.KeepAlive(msgType) // msgType is still "in use" until here
}

func fix() {
	msg := receiveMessage()
	msgType := make([]byte, 5)
	copy(msgType, msg)
	runtime.GC()
	printAlloc("fix")
	runtime.KeepAlive(msgType)
}
