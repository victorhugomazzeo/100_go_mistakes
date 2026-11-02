// Mistake #3: Misusing init functions
//
// init() runs before main, can't return an error (only panic) and needs
// global variables. Prefer a regular function that returns an error.
package main

import (
	"fmt"
	"strconv"
)

var port int

func init() {
	ports, err := strconv.Atoi("8080")
	if err != nil {
		panic(err)
	}
	port = ports
}

func main() {
	mistake()
	fix()
}

func mistake() {
	fmt.Println("[mistake] port:", port)
}

func fix() {
	ports, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("[fix] error:", err)
		return
	}

	fmt.Println("[fix] port:", ports)
}
