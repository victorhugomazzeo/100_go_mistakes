//go:build go1.21

// This file is compiled with the Go 1.21 loop rules (the build tag above), so
// range reuses a single variable for every iteration, as the book describes.
// That brings back the bug that Go 1.22 fixed. main.go uses the current rules.
package main

import "fmt"

func mistake() {
	var ptrs []*int
	for _, v := range []int{1, 2, 3} {
		ptrs = append(ptrs, &v) // same v every iteration: stores the same address 3 times
	}
	for _, p := range ptrs {
		fmt.Printf("mistake: %v\n", *p) // reads what is at each address: always the last value
	}
	fmt.Println()
}
