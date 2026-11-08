package customer

type Customer struct {
	id string
}

func New(id string) Customer {
	return Customer{id: id}
}

func (c Customer) ID() string {
	return c.id
}
