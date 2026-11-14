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

// loadOrderFix wraps the error with %w, so callers can still find ErrNotFound.
//
// Rule of thumb:
//   - %w: when the inner error is part of the contract, something callers
//     should be able to recognize (like ErrNotFound).
//   - %v: when the inner error is an implementation detail you don't want to
//     expose (like a database driver error); callers only get the text, for logs.
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
