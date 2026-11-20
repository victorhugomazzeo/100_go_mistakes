// Mistake #58: Not understanding race problems
//
// Data race: a variable shared by goroutines is written with no control (no
// lock) while others use it. Race condition: even with a lock, the result
// depends on which goroutine runs first.
package main

import (
	"fmt"
	"sync"
)

const goroutines = 1000
const increments = 1000

func main() {
	mistake1() // mistake of data race
	fix1()     // fix of data race

	mistake2() // mistake of race condition
	fix2()     // fix of race condition
}

func mistake1() {
	total := 0
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()

			for range increments {
				total++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("mistake: data race total:%v\n", total)
}

func fix1() {
	total := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()

			for range increments {
				mu.Lock()
				total++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("fix: data race total:%v\n", total)
}

func mistake2() {
	m := make(map[int]int)

	for range goroutines {
		i := 0
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			mu.Lock()
			i = 1
			mu.Unlock()
		}()
		go func() {
			defer wg.Done()

			mu.Lock()
			i = 2
			mu.Unlock()
		}()

		wg.Wait()
		m[i] += 1

	}

	fmt.Printf("mistake: race condition:%v\n", m)
}

func fix2() {
	m := make(map[int]int)

	for range goroutines {

		done := make(chan struct{})

		i := 0
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			i = 1
			close(done)
		}()
		go func() {
			defer wg.Done()

			<-done
			i = 2
		}()

		wg.Wait()
		m[i] += 1
	}

	fmt.Printf("fix: race condition:%v\n", m)
}
