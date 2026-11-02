// Mistake #4: Overusing getters and setters
//
// Getters and setters that just read/write a field add no value in Go:
// export the field. When one is useful, call it Name() (not GetName()).
package main

import "fmt"

type userWithGetters struct {
	name string
}

func (u *userWithGetters) GetName() string {
	return u.name
}

func (u *userWithGetters) SetName(name string) {
	u.name = name
}

type user struct {
	Name string
}

func mistake() {
	u := userWithGetters{}
	u.SetName("Gopher")
	fmt.Println("mistake:", u.GetName())
}

func main() {
	mistake()
	fix()
}

func fix() {
	u := user{Name: "Gopher"}
	fmt.Println("fix:", u.Name)
}
