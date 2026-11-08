// Mistake #22: Being confused about nil vs. empty slices
//
// A nil slice (var s []int) and an empty one ([]int{}) both have len 0, but
// only one is nil: encoding/json turns nil into null and empty into [].
package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	ID         string
	Operations []float32
}

func main() {
	mistake()
	fix()
}

func mistake() {
	var ops []float32
	b, _ := json.Marshal(customer{ID: "foo", Operations: ops})
	fmt.Println("mistake:", string(b))
}

func fix() {
	ops := make([]float32, 0)
	b, _ := json.Marshal(customer{ID: "foo", Operations: ops})
	fmt.Println("fix:", string(b))
}
