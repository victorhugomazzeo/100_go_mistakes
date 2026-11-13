// Mistake #42: Not knowing which type of receiver to use
//
// A value receiver works on a copy: changes made inside the method are lost.
// Use a pointer receiver to mutate (and when in doubt, choose the pointer).
package main

import "fmt"

type account struct {
	balance float64
}

func (a account) addByValue(v float64) {
	a.balance += v
}

func (a *account) addByPointer(v float64) {
	a.balance += v
}

func main() {
	mistake()
	fix()
}

func mistake() {
	acc := account{balance: 100}
	acc.addByValue(50)
	fmt.Println("mistake: balance:", acc.balance)
}

func fix() {
	acc := &account{balance: 100}
	acc.addByPointer(50)
	fmt.Println("fix: balance:", acc.balance)
}
