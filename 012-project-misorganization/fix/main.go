// Mistake #12: Project misorganization (the fix)
//
// One package per context (customer), with the type and its functions
// together, named after what it provides.
package main

import (
	"fmt"

	// One package per context: the Customer type and its functions live
	// together, so everything about customers is in one place.
	"github.com/victorhugomazzeo/100_go_mistakes/012-project-misorganization/fix/customer"
)

func main() {
	// The package name says what it provides: customer.All() reads as
	// "all customers".
	customers := customer.All()
	fmt.Printf("fix: %v\n", customers)
}
