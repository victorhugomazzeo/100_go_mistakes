// Mistake #59: Not understanding the concurrency impacts of a workload type
//
// CPU-bound work (computing) gains nothing from more goroutines than cores.
// I/O-bound work (waiting on an API, DB or disk) can use many more goroutines,
// up to what that external system can handle.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const jobs = 200 // how many tasks to process

func main() {
	mistake()
	fix()
}

func mistake() {
	cores := runtime.NumCPU()

	ioDuration := runPool(cores, ioTask)
	fmt.Printf("mistake: io, workers = cores(%v): %v\n", cores, ioDuration)

	cpuDuration := runPool(200, cpuTask)
	fmt.Printf("mistake: cpu, workers(200): %v\n", cpuDuration)
}

func fix() {
	cores := runtime.NumCPU()

	ioDuration := runPool(jobs, ioTask)
	fmt.Printf("fix: io, workers = jobs(%v): %v\n", jobs, ioDuration)

	cpuDuration := runPool(cores, cpuTask)
	fmt.Printf("fix: cpu, workers = cores(%v): %v\n", cores, cpuDuration)
}

// ioTask simulates I/O-bound work: 20ms waiting for an API or DB to answer.
func ioTask() {
	time.Sleep(20 * time.Millisecond)
}

// cpuTask simulates CPU-bound work: computing all the time.
func cpuTask() {
	sum := 0
	for i := range 30_000_000 {
		sum += i % 7
	}
	_ = sum
}

// runPool processes all jobs with a fixed number of workers and returns how long it took.
func runPool(workers int, task func()) time.Duration {
	start := time.Now()
	queue := make(chan struct{}) // the window where jobs arrive

	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range queue { // each worker takes jobs until the queue is closed
				task()
			}
		}()
	}

	for range jobs {
		queue <- struct{}{} // puts one job in the window
	}
	close(queue) // no more jobs: workers finish their loop
	wg.Wait()
	return time.Since(start)
}
