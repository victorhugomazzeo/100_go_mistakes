// Mistake #43: Never using named result parameters
//
// Name result parameters when it helps the reader: (lat, lng float32, err
// error) says what each value is. Use naked returns only in short functions.
package main

import "fmt"

func main() {
	mistake()
	fix()
}

func coordinatesUnnamed(address string) (float32, float32, error) {
	return -23.55, -46.63, nil //latitude, langitude
}

func coordinates(address string) (lat, lng float32, err error) {
	return -23.55, -46.63, nil
}

func mistake() {
	lng, lat, _ := coordinatesUnnamed("São Paulo") // guessed: longitude first?
	fmt.Println("mistake: lat:", lat, "lng:", lng)
}

func fix() {
	lat, lng, _ := coordinates("São Paulo")
	fmt.Println("fix: lat:", lat, "lng:", lng)
}
