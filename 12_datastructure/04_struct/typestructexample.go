package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) getFullName() string {
	return p.firstName + p.lastName
}

func main() {
	p1 := person{"James", "Bond", 22}
	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
	fmt.Println(p1.getFullName())
	fmt.Println(p2.getFullName())

}
