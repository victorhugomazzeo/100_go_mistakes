// Mistake #55: Mixing up concurrency and parallelism
//
// Concurrency is structure: tasks that can make progress independently.
// Parallelism is execution: tasks running at the same instant on many cores.
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

// work simulates a heavy task that only computes (no waiting).
func work() int {
	sum := 0
	for i := range 3_000_000_000 {
		sum += i % 7
	}
	return sum
}

// runTwoTasks runs work twice, concurrently, allowing Go to use `cores` cores.
func runTwoTasks(cores int) time.Duration {
	runtime.GOMAXPROCS(cores) // how many cooks Go may use at the same time
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(2)
	for range 2 {
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
	return time.Since(start)
}

func mistake() {
	start := time.Now()
	work()
	fmt.Println("mistake: one task alone:          ", time.Since(start))
	fmt.Println("mistake: two concurrent tasks, 1 core:", runTwoTasks(1))
}

func fix() {
	fmt.Println("fix: two concurrent tasks, all cores:", runTwoTasks(runtime.NumCPU()))
}
