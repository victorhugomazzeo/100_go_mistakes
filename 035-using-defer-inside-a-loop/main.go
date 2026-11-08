// Mistake #35: Using defer inside a loop
//
// defer runs when the surrounding function returns, not at the end of each
// loop iteration. Move the loop body into a function so defer runs each time.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	for i := 1; i <= 3; i++ {
		fmt.Println("mistake: open", i)
		defer fmt.Println("mistake: close", i)
	}
	fmt.Println("mistake: loop done")
}

func fix() {
	for i := 1; i <= 3; i++ {
		process(i) // each call is a whole function: its defer runs when it returns
	}
	fmt.Println("fix: loop done")
}

func process(i int) {
	fmt.Println("fix: open", i)
	defer fmt.Println("fix: close", i)
	// work with the "file" here
}
