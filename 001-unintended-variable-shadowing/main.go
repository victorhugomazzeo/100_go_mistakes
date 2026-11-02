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
	mistake()
	fix()
}

func mistake() {
	var n int

	if true {
		n, err := strconv.Atoi("42") // := declares a NEW n (and err) inside the if
		fmt.Println("[mistake] dentro do if:", n, err)
	}

	fmt.Println("[mistake] fora do if:", n) // 0
}

func fix() {
	var n int
	var err error

	if true {
		n, err = strconv.Atoi("42") // = assigns to the outer n and err
		fmt.Println("[fix] dentro do if:", n, err)
	}

	fmt.Println("[fix] fora do if:", n) // 42
}
