// Mistake #54: Not handling defer errors
//
// defer r.Close() drops the error Close returns. Ignore it explicitly
// (_ = r.Close()) or propagate it through a named result set by the defer.
package main

import (
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

// resource simulates a file or DB connection whose Close can fail.
type resource struct{}

func (r resource) Close() error {
	return errors.New("close failed: data not flushed")
}

func open() (resource, error) {
	return resource{}, nil
}

func saveMistake() error {
	r, err := open()
	if err != nil {
		return err
	}
	defer r.Close() // the error from Close is silently dropped
	// ... write data ...
	return nil
}

func mistake() {
	fmt.Println("mistake:", saveMistake())
}

func saveFix() (err error) { // named result: the defer can change it
	r, err := open()
	if err != nil {
		return err
	}
	defer func() { err = errors.Join(err, r.Close()) }()
	// ... write data ...
	return nil
}

func fix() {
	fmt.Println("fix:", saveFix())
}
