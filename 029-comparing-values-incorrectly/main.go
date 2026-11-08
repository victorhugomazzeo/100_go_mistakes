// Mistake #29: Comparing values incorrectly
//
// == only works on comparable types, and a struct with a slice field isn't.
// Compare it with your own equal method, or reflect.DeepEqual (slower).
package main

import (
	"fmt"
	"reflect"
	"slices"
)

// customer1 has only comparable fields, so == works on it.
type customer1 struct {
	id string
}

// customer2 has a slice field, so == doesn't compile on it.
type customer2 struct {
	id         string
	operations []float64
}

func main() {
	mistake()
	fix()
}

func mistake() {
	a := customer1{id: "x"}
	b := customer1{id: "x"}
	fmt.Println("mistake: customer1 == customer1:", a == b) // fine: only a string inside

	c1 := customer2{id: "x", operations: []float64{1}}
	c2 := customer2{id: "x", operations: []float64{1}}
	// fmt.Println(c1 == c2) // doesn't compile: struct containing []float64 cannot be compared
	_, _ = c1, c2
}

func fix() {
	c1 := customer2{id: "x", operations: []float64{1}}
	c2 := customer2{id: "x", operations: []float64{1}}
	c3 := customer2{id: "x", operations: []float64{2}}

	fmt.Println("fix: c1.equal(c2):", c1.equal(c2))                           // same id and operations
	fmt.Println("fix: c1.equal(c3):", c1.equal(c3))                           // operations differ
	fmt.Println("fix: reflect.DeepEqual(c1, c2):", reflect.DeepEqual(c1, c2)) // works too, but slower
}

// equal defines what "equal" means for customer2, since == can't.
func (a customer2) equal(b customer2) bool {
	return a.id == b.id && slices.Equal(a.operations, b.operations)
}
