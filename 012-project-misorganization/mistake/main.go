// Mistake #12: Project misorganization (the mistake)
//
// One type and one function split by layer into two tiny packages, named
// after what they contain (models, services) instead of what they provide.
package main

import (
	"fmt"

	// services has only the functions; the Customer type lives in models.
	// To understand "customer" you have to open two packages.
	"github.com/victorhugomazzeo/100_go_mistakes/012-project-misorganization/mistake/services"
)

func main() {
	// The package name says what it contains (services), not what it
	// provides: services.Customers() doesn't read as a sentence.
	fmt.Println("mistake:", services.Customers())
}
