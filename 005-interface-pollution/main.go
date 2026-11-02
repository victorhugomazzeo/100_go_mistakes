// Mistake #5: Interface pollution
//
// Don't create an interface before you need it (e.g., only one implementation).
// Go interfaces are implicit: add one later, when a real need shows up.
package main

import "fmt"

type emailNotifier struct{}

func (e *emailNotifier) Notify(msg string) {
	fmt.Println("email:", msg)
}

type Notifier interface {
	Notify(msg string)
}

func mistake() {
	var n Notifier = &emailNotifier{}
	n.Notify("mistake")
}

func fix() {
	n := emailNotifier{}
	n.Notify("fix")
}

func main() {
	mistake()
	fix()
}
