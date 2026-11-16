// Mistake #56: Thinking concurrency is always faster
//
// Goroutines aren't free: creating and scheduling one costs more than a tiny
// task. Give each goroutine a big chunk of work, or just stay sequential.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	mistake()
	fix()
}

const n = 1_000_000

// square is the tiny work done on each item.
func square(x int) int {
	return x * x
}

// newSlice returns 0, 1, 2, ..., n-1.
func newSlice() []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

func mistake() {
	s := newSlice()
	start := time.Now()

	for i := range s {
		s[i] = square(s[i])
	}
	fmt.Println("mistake: sequential:            ", time.Since(start))

	s = newSlice()
	start = time.Now()
	var wg sync.WaitGroup
	for i := range s {
		wg.Add(1) // one more task pending, for each item
		go func() {
			defer wg.Done()
			s[i] = square(s[i]) // each goroutine touches only its own slot
		}()
	}
	wg.Wait()
	fmt.Println("mistake: one goroutine per item:", time.Since(start))
}

func fix() {
	s := newSlice()
	start := time.Now()

	workers := runtime.NumCPU() // one goroutine per core: 20 on your machine
	size := len(s) / workers    // items per goroutine: 50,000

	var wg sync.WaitGroup
	for w := range workers {
		from := w * size
		to := from + size
		if w == workers-1 {
			to = len(s) // the last goroutine also takes any leftover items
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := from; i < to; i++ { // a big chunk of work per goroutine
				s[i] = square(s[i])
			}
		}()
	}
	wg.Wait()
	fmt.Println("fix: one goroutine per core:", time.Since(start))
}
