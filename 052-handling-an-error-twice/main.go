// Mistake #52: Handling an error twice
//
// Logging an error is handling it: either log it or return it, never both, or
// one failure fills the log many times. Wrap it with context and log it once.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)          // no date/time prefix
	log.SetOutput(os.Stdout) // same output as fmt, so lines stay in order
	mistake()
	fix()
}

func validateLatMistake(lat float32) error {
	if lat > 90 || lat < -90 {
		log.Printf("invalid latitude: %.2f", lat)        // logs...
		return fmt.Errorf("invalid latitude: %.2f", lat) // ...and returns
	}
	return nil
}

func getRouteMistake(srcLat, dstLat float32) error {
	if err := validateLatMistake(srcLat); err != nil {
		log.Println("failed to validate source") // logs again...
		return err                               // ...and returns again
	}
	return nil
}

func mistake() {
	if err := getRouteMistake(200, 10); err != nil { // 200 is an invalid latitude
		log.Println("mistake:", err)
	}
}

func validateLatFix(lat float32) error {
	if lat > 90 || lat < -90 {
		return fmt.Errorf("invalid latitude: %.2f", lat) // only returns
	}
	return nil
}

func getRouteFix(srcLat, dstLat float32) error {
	if err := validateLatFix(srcLat); err != nil {
		return fmt.Errorf("validate source: %w", err) // adds context, still only returns
	}
	return nil
}

func fix() {
	if err := getRouteFix(200, 10); err != nil {
		log.Println("fix:", err) // the only place that logs
	}
}
