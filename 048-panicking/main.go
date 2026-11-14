// Mistake #48: Panicking
//
// panic stops the program: keep it for truly unrecoverable cases (a bug, a
// missing mandatory dependency). For expected failures, return an error.
package main

import (
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

func checkAgeMistake(age int) {
	if age < 0 {
		panic("invalid age") // stops everything for an ordinary validation
	}
}

func mistake() {
	defer func() {
		if r := recover(); r != nil { // catches the panic so the program can go on
			fmt.Println("mistake: the program would crash here:", r)
		}
	}()
	checkAgeMistake(-1)
	fmt.Println("mistake: this line never runs")
}

func checkAgeFix(age int) error {
	if age < 0 {
		return errors.New("invalid age") // an ordinary failure: just report it
	}
	return nil
}

func fix() {
	if err := checkAgeFix(-1); err != nil {
		fmt.Println("fix: invalid input, handled:", err)
		return
	}

	fmt.Println("fix: age is valid")
}
