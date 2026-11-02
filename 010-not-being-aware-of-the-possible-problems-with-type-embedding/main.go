// Mistake #10: Not being aware of the possible problems with type embedding
//
// Embedding promotes every field and method of the embedded type, so they
// become part of your type's API. Embed only what callers should use.
package main

import (
	"fmt"
	"sync"
)

type counterEmbedded struct {
	sync.Mutex
	n int
}

type counterNotEmbedded struct {
	n  int
	mu sync.Mutex
}

func (c *counterNotEmbedded) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func main() {
	mistake()
	fix()
}

func mistake() {
	c := counterEmbedded{}
	c.Lock()
	c.n++
	c.Unlock()
	fmt.Printf("mistake: %d\n", c.n)
}

func fix() {
	c := counterNotEmbedded{}
	c.inc()
	fmt.Printf("fix: %d\n", c.n)
}
