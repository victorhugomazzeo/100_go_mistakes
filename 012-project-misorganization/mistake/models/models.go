// Package models holds only types: a layer, named after what it contains.
package models

// Customer and Name must be exported: services, another package, builds it.
type Customer struct {
	Name string
}
