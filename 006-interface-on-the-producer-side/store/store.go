// Package store is the producer: it has the implementation.
package store

type Store struct{}

func (s Store) Save(name string) {
}

func (s Store) GetAll() []string {
	return []string{"Gopher"}
}

type CustomerStorage interface {
	Save(name string)
	GetAll() []string
}
