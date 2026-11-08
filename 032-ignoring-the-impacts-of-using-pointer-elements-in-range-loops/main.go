// Mistake #32: Ignoring the impacts of using pointer elements in range loops
//
// Before Go 1.22, a range loop reused one variable v for every iteration, so
// &v stored the same address each time. Since 1.22 each iteration gets its own.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func fix() {
	var ptrs []*int
	for _, v := range []int{1, 2, 3} {
		ptrs = append(ptrs, &v) // a new v every iteration: stores 3 different addresses
	}
	for _, p := range ptrs {
		fmt.Printf("fix: %v\n", *p) // reads what is at each address: 1, 2, 3
	}
	fmt.Println()
}
