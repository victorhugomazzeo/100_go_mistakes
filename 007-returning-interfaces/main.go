// Mistake #7: Returning interfaces
//
// Return concrete types, accept interfaces. Returning an interface hides
// the rest of the type and forces one abstraction on every caller.
package main

import "fmt"

type store struct{}

func (s store) GetAll() []string {
	return []string{"Gopher"}
}

func (s store) Count() int {
	return 1
}

type customerGetter interface {
	GetAll() []string
}

func newStoreAsInterface() customerGetter {
	return store{}
}

func newStoreAsConcrete() store {
	return store{}
}

func main() {
	mistake()
	fix()
}

func mistake() {
	s := newStoreAsInterface()
	fmt.Println("mistake:", s.GetAll())
}

func fix() {
	s := newStoreAsConcrete()
	fmt.Println("fix:", s.GetAll(), s.Count())
}
