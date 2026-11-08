// Mistake #19: Not understanding floating-points
//
// Floats are stored in binary, where 0.1 and 0.2 are repeating fractions
// (like 1/3 in decimal) that get rounded: 0.1 + 0.2 is 0.30000000000000004.
// Never compare floats with ==; check their difference is within a delta.
package main

import (
	"math"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	a, b := 0.1, 0.2
	if a+b == 0.3 {
		println("mistake: a + b is 0.3")
	} else {
		println("mistake: a + b is NOT 0.3. actually:", a+b)
	}
}

func fix() {
	a, b := 0.1, 0.2
	if almostEqual(a+b, 0.3) {
		println("fix: a + b is 0.3")
	} else {
		println("fix: a + b is NOT 0.3. actually:", a+b)
	}
}

func almostEqual(x, y float64) bool {
	return math.Abs(x-y) < 1e-9
}
