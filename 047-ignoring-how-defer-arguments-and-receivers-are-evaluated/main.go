// Mistake #47: Ignoring how defer arguments and receivers are evaluated
//
// defer evaluates its arguments right away, not when it runs: defer f(status)
// keeps the current status. Wrap the call in a closure to read it at the end.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	status := "pending"
	defer fmt.Println("mistake: status at the end:", status)

	// ... the work happens here ...
	status = "success"
}

func fix() {
	status := "pending"
	defer func() {
		fmt.Println("fix: status at the end:", status) // reads status when it runs
	}()
	status = "success"
}
