// Mistake #6: Interface on the producer side
//
// Interfaces belong to the consumer, not the producer: let the code that
// uses an abstraction define it, with only the methods it needs.
package main

import (
	"fmt"

	"github.com/victorhugomazzeo/100_go_mistakes/006-interface-on-the-producer-side/store"
)

type customerGetter interface {
	GetAll() []string
}

func main() {
	mistake()
	fix()
}

func mistake() {
	var s store.CustomerStorage = store.Store{}
	fmt.Println("mistake:", s.GetAll())
}

func fix() {
	var g customerGetter = store.Store{}
	fmt.Println("fix:", g.GetAll())
}
