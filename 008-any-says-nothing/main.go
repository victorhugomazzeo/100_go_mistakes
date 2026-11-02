// Mistake #8: any says nothing
//
// any accepts every type, so it tells callers nothing and the compiler
// can't check it: mistakes only show up at runtime. Prefer precise types.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func doubleAny(v any) int {
	return v.(int) * 2
}

func double(n int) int {
	return n * 2
}

func mistake() {
	fmt.Println("mistake:", doubleAny(21))
}

func fix() {
	fmt.Println("fix:", double(21))
}
