package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//this does not work, as they are different type, one is type "foo", another one is type int
	//fmt.Println myAge + yourAge)
	//Error, mismatch type foo and int

	//this conversion works:
	fmt.Println(int(myAge) + yourAge)
}
