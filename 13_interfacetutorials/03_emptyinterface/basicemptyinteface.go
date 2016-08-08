package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

//You can put anything in empty interface, and use it with cautions cause Go is strictly typed
func printIt(a interface{}) {
	fmt.Println(a)
}
func main() {
	bill := dog{animal{"gaugau"}, true}
	conmeo := cat{animal{"meomeo"}, false}

	printIt(bill)
	printIt(conmeo)

	//// We can have a slice of empty interface
	conmeo2 := cat{animal{"meomeomeomeo"}, true}
	conbill2 := dog{animal{"gaugaugaugau"}, true}

	// Here is the slice of empty interface
	animalList := []interface{}{bill, conmeo, conmeo2, conbill2}

	fmt.Println(animalList)
}
