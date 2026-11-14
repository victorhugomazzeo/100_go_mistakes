// Mistake #50: Comparing an error type inaccurately
//
// A type assertion only checks the outer error, so it fails once the error is
// wrapped with %w. Use errors.As to search the whole chain for a type.
package main

import (
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

// transientError is a temporary failure: retrying may work.
type transientError struct {
	err error
}

func (t transientError) Error() string {
	return "transient error: " + t.err.Error()
}

// queryDB simulates a dropped database connection.
func queryDB(id string) error {
	return transientError{err: errors.New("connection reset")}
}

// getAmount adds context with %w, as the #49 taught.
func getAmount(id string) error {
	if err := queryDB(id); err != nil {
		return fmt.Errorf("get amount %s: %w", id, err)
	}
	return nil
}

func mistake() {
	err := getAmount("12345")
	_, ok := err.(transientError) // type assertion: checks only the outer error
	fmt.Printf("mistake: is transient? %t (actual type: %T)\n", ok, err)
}

func fix() {
	err := getAmount("12345")
	var te transientError     // where errors.As puts what it finds
	ok := errors.As(err, &te) // searches the whole chain for a transientError
	fmt.Println("fix: is transient?", ok, "| cause:", te.err)
}
