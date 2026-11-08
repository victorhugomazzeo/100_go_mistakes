// Mistake #31: Ignoring how arguments are evaluated in range loops
//
// range evaluates its expression once, before the loop, and iterates over a
// copy: with an array, changes made inside the loop don't reach the value v.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println("mistake:", v)
		}
	}
}

func fix() {
	a := [3]int{0, 1, 2}
	for i := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println("fix:", a[2])
		}
	}

}
