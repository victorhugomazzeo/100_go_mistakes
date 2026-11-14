// Mistake #51: Comparing an error value inaccurately
//
// Comparing a sentinel error with == fails once it is wrapped with %w. Use
// errors.Is, which unwraps the chain: errors.Is(err, sql.ErrNoRows).
package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

// query simulates a SELECT that finds nothing, adding context with %w.
func query(id int) error {
	return fmt.Errorf("query customer %d: %w", id, sql.ErrNoRows)
}

func mistake() {
	err := query(42)
	fmt.Println("mistake: no rows?", err == sql.ErrNoRows)
}

func fix() {
	err := query(42)
	fmt.Println("fix: no rows?", errors.Is(err, sql.ErrNoRows))
}
