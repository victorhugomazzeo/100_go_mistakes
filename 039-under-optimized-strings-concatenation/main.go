// Mistake #39: Under-optimized strings concatenation
//
// Strings are immutable: s += v in a loop builds a brand-new string every
// time. Join many strings with strings.Builder (Grow it if the size is known).
package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	mistake()
	fix()
}

func mallocs() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Mallocs
}

func input() []string {
	values := make([]string, 1000)
	for i := range values {
		values[i] = strings.Repeat("x", 1000) // "xxx...x", 1000 times
	}
	return values
}

func mistake() {
	values := input()
	before, start := mallocs(), time.Now()

	s := ""
	for _, v := range values {
		s += v
	}

	fmt.Println("mistake:", time.Since(start), mallocs()-before, "allocations")
}

func fix() {
	values := input()
	before, start := mallocs(), time.Now()

	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(v)
	}

	s := sb.String()
	fmt.Println("fix:", time.Since(start), mallocs()-before, "allocations", len(s))
}
