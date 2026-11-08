// Mistake #17: Creating confusion with octal literals
//
// An integer literal starting with 0 is octal: 010 is 8, not 10. Write octal
// explicitly with 0o (0o10), as in file permissions like 0o644.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	sum := 100 + 010                             // 010 is octal for 8, not decimal 10
	fmt.Printf("mistake: 100 + 010 = %d\n", sum) // prints 108, not 110
}

func fix() {
	sum := 100 + 0o10                         // 0o10 is octal for 8, explicitly written
	fmt.Printf("fix: 100 + 0o10 = %d\n", sum) // prints 108, not 110
}
