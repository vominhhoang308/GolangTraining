package main

import "fmt"

func main() {
	var smallnumber int
	var largenumber int
	fmt.Println("please enter one small number and one large number, we can give the remainder")

	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallnumber)

	fmt.Print("Please enter a large number: ")
	fmt.Scan(&largenumber)

	fmt.Println("The remainder for ", largenumber, "/", smallnumber, " is ", largenumber%smallnumber)
}
