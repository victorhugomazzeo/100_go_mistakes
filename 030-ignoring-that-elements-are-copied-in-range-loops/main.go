// Mistake #30: Ignoring that elements are copied in range loops
//
// In for _, v := range s, v is a copy of each element: changing v doesn't
// change s. To modify the elements, go through the index: s[i].field = ...
package main

import "fmt"

func main() {
	mistake()
	fix()
}

type account struct {
	balance float32
}

func createAccounts() []account {
	return []account{{balance: 100}, {balance: 200}, {balance: 300}}
}

func mistake() {
	accounts := createAccounts()
	for _, a := range accounts {
		a.balance += 1000
	}
	fmt.Println("mistake:", accounts)
}

func fix() {
	accounts := createAccounts()
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println("fix:", accounts)
}
