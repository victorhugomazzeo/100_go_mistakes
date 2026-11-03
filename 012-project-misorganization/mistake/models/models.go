// Package models holds only types: a layer, named after what it contains.
package models

// Customer must be exported: services, another package, needs it.
type Customer struct {
	Name string
}
