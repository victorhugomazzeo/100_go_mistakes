// Mistake #11: Not using the functional options pattern
//
// For optional config, let callers pass only what they change: each option is
// a function that edits the struct, applied over defaults in the constructor.
package main

import "fmt"

type server struct {
	port    int
	timeout int
}

type option func(*server)

func main() {
	mistake()
	fix()
}

func newServerAllParams(port, timeout int) *server {
	return &server{
		port:    port,
		timeout: timeout,
	}
}

func withPort(port int) option {
	return func(s *server) {
		s.port = port
	}
}

func newServer(opts ...option) server {
	s := server{
		port:    8080,
		timeout: 30,
	}
	for _, opt := range opts {
		opt(&s)
	}
	return s
}

func mistake() {
	s := newServerAllParams(9090, 30) // caller must pass all params, even if they don't want to change them
	fmt.Printf("mistake: %v\n", s)
}

func fix() {
	s := newServer(withPort(9090)) // caller only passes what they want to change
	fmt.Printf("fix: %v\n", s)
}
