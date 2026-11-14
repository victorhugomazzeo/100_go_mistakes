// Mistake #45: Returning a nil receiver
//
// An interface is nil only when both its type and value are nil: returning a
// nil *T as an error gives a non-nil error. Return an explicit nil instead.
package main

import (
	"fmt"
	"strings"
)

type validationError struct {
	msgs []string
}

func (v *validationError) Error() string {
	return strings.Join(v.msgs, "; ")
}

func main() {
	mistake()
	fix()
}

func validateMistake(name string) error {
	var verr *validationError // starts as a nil pointer
	if name == "" {
		verr = &validationError{msgs: []string{"name is empty"}}
	}

	return verr // nil pointer, but returned as error
}

func mistake() {
	err := validateMistake("John")
	fmt.Printf("mistake: err == nil? %t, type: %T\n", err == nil, err)
}

func validateFix(name string) error {
	var err error // the box itself: starts empty, like "var s speaker"
	if name == "" {
		err = &validationError{msgs: []string{"name is empty"}}
	}
	return err // still the empty box when the name is valid
}

func fix() {
	err := validateFix("John")
	fmt.Printf("fix: err == nil? %t, type: %T\n", err == nil, err)
}
