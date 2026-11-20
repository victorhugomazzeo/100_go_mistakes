// Mistake #61: Propagating an inappropriate context
//
// The server cancels a request's context as soon as the response is sent. If
// a goroutine that keeps running after that (like publishing a message) uses
// the same context, it gets canceled midway. Detach it: context.WithoutCancel.
package main

import (
	"context"
	"fmt"
	"time"
)

// key is the type for context keys (a custom type avoids clashes with other packages).
type key string

func main() {
	mistake()
	fix()
}

func mistake() {
	ctx := context.WithValue(context.Background(), key("requestID"), "req-42")
	reqCtx, cancel := context.WithCancel(ctx)

	result := make(chan error)

	go func() {
		result <- publish(reqCtx, "order 42 paid")
	}()

	cancel()

	fmt.Printf("mistake: %v\n", <-result)
}

func fix() {
	ctx := context.WithValue(context.Background(), key("requestID"), "req-42")
	reqCtx, cancel := context.WithCancel(ctx)

	result := make(chan error)

	go func() {
		result <- publish(context.WithoutCancel(reqCtx), "order 42 paid")
	}()

	cancel()

	fmt.Printf("fix: %v\n", <-result)
}

// publish simulates sending a message to another system: it takes 100ms
// and gives up if the context is done first.
func publish(ctx context.Context, msg string) error {
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("  published %q (request id %v)\n", msg, ctx.Value(key("requestID")))
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
