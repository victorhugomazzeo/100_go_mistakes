// Mistake #63: Not being careful with goroutines and loop variables
//
// Before Go 1.22, goroutines started in a loop all shared the same loop
// variable, so they often printed the last value instead of their own.
package main

import (
	"fmt"
	"sync"
)

func main() {
	mistake()
	fix()
}

func fix() {
	var wg sync.WaitGroup
	items := []int{1, 2, 3}

	fmt.Print("fix:")

	for _, i := range items {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Print(i)
		}()
	}

	wg.Wait()
	fmt.Println()
}
