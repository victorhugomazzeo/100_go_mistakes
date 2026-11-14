// Mistake #53: Not handling an error
//
// Dropping an error silently looks like a bug to the next reader. If ignoring
// it is on purpose, say so: _ = notify(), with a comment explaining why.
package main

import (
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

func notify() error {
	return errors.New("failed to notify")
}

func mistake() {
	fmt.Println("mistake: order saved")
	notify()
}

func fix() {
	fmt.Println("fix: order saved")

	// Notifications are best effort: a failed e-mail must not undo a saved order.
	_ = notify()
}
