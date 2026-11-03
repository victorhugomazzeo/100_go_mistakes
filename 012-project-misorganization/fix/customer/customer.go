package customer

type Customer struct {
	Name string
}

func All() []Customer {
	return []Customer{
		{Name: "Gopher"},
	}
}
