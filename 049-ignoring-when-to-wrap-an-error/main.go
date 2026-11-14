// Mistake #49: Ignoring when to wrap an error
//
// Wrap with %w to add context and keep the source error reachable through
// errors.Is; %v only copies its text. Don't wrap if callers shouldn't see it.
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is returned when a record doesn't exist.
var ErrNotFound = errors.New("not found")

func main() {
	mistake()
	fix()
}

func findUser(id int) error {
	return ErrNotFound // simulates: no user with this id
}

func loadOrderMistake(id int) error {
	if err := findUser(id); err != nil {
		return fmt.Errorf("load order %d: %v", id, err) // %v: copies only the text
	}
	return nil
}

func mistake() {
	err := loadOrderMistake(42)
	fmt.Println("mistake:", err, "| is not found?", errors.Is(err, ErrNotFound))
}

func loadOrderFix(id int) error {
	if err := findUser(id); err != nil {
		return fmt.Errorf("load order %d: %w", id, err) // %w: keeps the original inside
	}
	return nil
}

func fix() {
	err := loadOrderFix(42)
	fmt.Println("fix:", err, "| is not found?", errors.Is(err, ErrNotFound))
}
