// Mistake #18: Neglecting integer overflows
//
// At run time integer overflows are silent: MaxInt32 + 1 wraps to MinInt32,
// no panic. Check the bounds yourself before the operation when it matters.
package main

import (
	"fmt"
	"math"
)

func main() {
	mistake()
	fix()
}

func mistake() {

	counter := int32(math.MaxInt32)
	fmt.Printf("mistake: before counter: %d\n", counter)
	counter++
	fmt.Printf("mistake: after counter: %d\n", counter)

	// var c int32 = math.MaxInt32 + 1
	// This will not compile, because the constant expression overflows int32.
}

func fix() {
	c, err := inc32(math.MaxInt32)
	if err != nil {
		fmt.Println("fix:", err)
		return
	}
	fmt.Println("fix: after counter:", c)
}

func inc32(counter int32) (int32, error) {
	if counter == math.MaxInt32 {
		return 0, fmt.Errorf("int32 overflow: cannot increment %d", counter)
	}
	return counter + 1, nil
}
