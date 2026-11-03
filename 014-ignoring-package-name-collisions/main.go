// Mistake #14: Ignoring package name collisions
//
// A variable named like a package (redis := redis.NewClient()) hides the
// package for the rest of the scope: pick another name or alias the import.
package main

import (
	"fmt"

	"github.com/victorhugomazzeo/100_go_mistakes/014-ignoring-package-name-collisions/redis"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	redis := redis.NewClient()
	fmt.Println("mistake:", redis.Get("foo"))

	// redis = redis.NewClient() // error: redis package is hidden by the variable named redis
}

func fix() {
	client := redis.NewClient()
	fmt.Println("fix:", client.Get("foo"))
}
