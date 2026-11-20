// Mistake #60: Misunderstanding Go contexts
//
// A context travels with a call that may take long (an API, a database). It
// carries a deadline, a cancellation signal and values, so the caller can stop
// a slow call instead of waiting forever.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mistake()
	fix()
}

func mistake() {
	now := time.Now()
	ctx := context.Background()
	v, err := slowAPI(ctx)
	fmt.Printf("mistake: value:%v err:%v duration:%v\n", v, err, time.Since(now))
}

// slowAPI simulates an external API that takes 3 seconds to answer.
// It gives up early if the context is done (deadline reached or canceled).
func slowAPI(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second): // the answer arrives after 3s
		return "data", nil
	case <-ctx.Done(): // the caller's deadline was reached first
		return "", ctx.Err()
	}
}

func fix() {
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	v, err := slowAPI(ctx)
	fmt.Printf("fix: value:%v err:%v duration:%v\n", v, err, time.Since(now))
}
