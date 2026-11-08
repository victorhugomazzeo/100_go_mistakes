// Package customer provides a simple representation of a customer with an ID.
package customer

// Customer represents a customer with an ID.
type Customer struct {
	id string
}

// New creates a new Customer with the given ID.
func New(id string) Customer {
	return Customer{id: id}
}

// ID returns the ID of the customer.
func (c Customer) ID() string {
	return c.id
}
