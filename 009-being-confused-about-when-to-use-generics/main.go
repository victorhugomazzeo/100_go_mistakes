// Mistake #9: Being confused about when to use generics
//
// Use generics when the same logic repeats for different types. Don't use
// them when a plain interface or a concrete type already does the job.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func mistake() {
	fmt.Println("mistake:", sumInts([]int{1, 2, 3}))
	fmt.Println("mistake:", sumFloats([]float64{1.0, 2.0, 3.0}))
}

func sumInts(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func sumFloats(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

func sum[T int | float64](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func fix() {
	fmt.Println("fix:", sum([]int{1, 2, 3}))
	fmt.Println("fix:", sum([]float64{1.0, 2.0, 3.0}))
}
