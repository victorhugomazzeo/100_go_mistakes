// Mistake #1: Unintended variable shadowing
//
// Inside a block (if, for...), := declares a NEW variable that hides
// the outer one with the same name. The outer one never changes.
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int

	if true {
		n, _ := strconv.Atoi("42")
		fmt.Println("dentro do if:", n)
	}

	fmt.Println("fora do if:", n)
}
