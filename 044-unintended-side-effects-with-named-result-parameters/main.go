// Mistake #44: Unintended side effects with named result parameters
//
// Named results start at their zero value: returning err before assigning it
// returns nil and hides the failure. Assign the error before returning it.
package main

import (
	"errors"
	"fmt"
)

func main() {
	mistake()
	fix()
}

// checkService simulates an external service that is down.
func checkService() error {
	return errors.New("service unavailable")
}

func coordinatesMistake(address string) (lat, lng float32, err error) {
	if checkService() != nil {
		return 0, 0, err // err was never assigned: it is still nil
	}

	return -23.55, -46.63, nil
}
func mistake() {
	lat, lng, err := coordinatesMistake("São Paulo")
	fmt.Println("mistake:", lat, lng, "err:", err)
}

func fix() {
	lat, lng, err := coordinatesFix("São Paulo")
	fmt.Println("fix:", lat, lng, "err:", err)
}

func coordinatesFix(address string) (lat, lng float32, err error) {
	if err = checkService(); err != nil {
		return 0, 0, err
	}
	return -23.55, -46.63, nil
}
