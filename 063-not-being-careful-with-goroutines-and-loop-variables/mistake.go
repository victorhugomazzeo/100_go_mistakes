//go:build go1.21

// This file is compiled with the Go 1.21 loop rules (the build tag above), so
// the loop reuses one variable for every iteration and the goroutines share it.
// That brings back the bug that Go 1.22 fixed. main.go uses the current rules.
package main

import (
	"fmt"
	"sync"
)

func mistake() {
	var wg sync.WaitGroup
	items := []int{1, 2, 3}

	fmt.Print("mistake:")

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
