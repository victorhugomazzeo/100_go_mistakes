// Mistake #34: Ignoring how the break statement works
//
// break exits the innermost for, switch or select: inside a switch it leaves
// the switch, not the loop. Use a label (break loop) to exit the loop.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	for i := 0; i < 5; i++ {
		fmt.Printf("mistake: %d ", i)

		switch i {
		default:
		case 2:
			break
		}
	}
	fmt.Println()
}

func fix() {
loop: // the label names the for below
	for i := 0; i < 5; i++ {
		fmt.Printf("fix: %d ", i)

		switch i {
		default:
		case 2:
			break loop // exits the for named loop, not the switch
		}
	}
	fmt.Println()
}
