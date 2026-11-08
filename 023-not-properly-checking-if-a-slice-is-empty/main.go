// Mistake #23: Not properly checking if a slice is empty
//
// Check emptiness with len(s) == 0, not s == nil: len covers both nil and
// empty slices (and maps). Callers shouldn't care which one they got.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	ops := gotOperations("")
	if ops != nil {
		fmt.Println("mistake: handling", ops)
	}
}

func fix() {
	ops := gotOperations("")
	if len(ops) != 0 {
		fmt.Println("fix: handling", ops)
	} else {
		fmt.Println("fix: nothing to handle")
	}

}

func gotOperations(id string) []float32 {
	if id == "" {
		return make([]float32, 0)
	}
	return []float32{10, 20}
}
