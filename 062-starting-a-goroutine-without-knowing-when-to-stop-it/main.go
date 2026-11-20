// Mistake #62: Starting a goroutine without knowing when to stop it
//
// Every goroutine needs a plan to stop. One that loops forever with no way to
// stop it leaks; give it a close method that signals it and waits for it.
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type watcher struct {
	cancel func()
	done   chan struct{}
}

func main() {
	mistake()
	fix()
}

func startWatcherMistake() {
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
		}
	}()
}

func mistake() {
	before := runtime.NumGoroutine()
	startWatcherMistake()
	time.Sleep(120 * time.Millisecond)
	after := runtime.NumGoroutine()
	fmt.Printf("mistake: goroutines before:%v after:%v\n", before, after)
}

func fix() {
	before := runtime.NumGoroutine()
	w := newWatcher()
	time.Sleep(120 * time.Millisecond)
	w.close()
	after := runtime.NumGoroutine()
	fmt.Printf("fix: goroutines before:%v after:%v\n", before, after)
}

func newWatcher() *watcher {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{cancel: cancel, done: make(chan struct{})}
	go func() {
		defer close(w.done) // closes the channel: "I'm done", wakes up whoever waits on <-w.done

		for {
			select {
			case <-time.After(50 * time.Millisecond):

			case <-ctx.Done():
				fmt.Printf("fix: watcher: resources closed\n")
				return
			}
		}
	}()

	return w
}

func (w *watcher) close() {
	w.cancel()
	<-w.done
}
