// Mistake #57: Being puzzled about when to use channels or mutexes
//
// They do different jobs: a mutex protects state that goroutines share; a
// channel passes data or signals between goroutines. Don't force channels.
package main

import (
	"fmt"
	"sync"
	"time"
)

const goroutines = 1000
const incrementsEach = 1000

func main() {
	mistake()
	fix()
}

func mistake() {
	start := time.Now()

	increments := make(chan int) // the window where the +1s arrive
	result := make(chan int)     // the window where the final total comes back

	// 1. The goroutine that owns the counter: the only one that adds.
	go func() {
		total := 0
		for v := range increments { // receives every +1 until the channel is closed
			total += v
		}
		result <- total // sends the final total back
	}()

	// 2. A thousand goroutines sending +1 through the channel.
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			for range incrementsEach {
				increments <- 1 // instead of adding, each one sends a +1
			}
		}()
	}

	// 3. When all senders are done, close the channel and get the total.
	wg.Wait()
	close(increments) // tells the owner: no more +1s, finish the range
	total := <-result // waits for the owner to send the total

	fmt.Println("mistake (channel):", total, time.Since(start))
}

func fix() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(goroutines)

	var mu sync.Mutex

	total := 0

	for range goroutines {
		go func() {
			defer wg.Done()
			for range incrementsEach {
				mu.Lock()
				total++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("fix (mutex):", total, time.Since(start))
}
