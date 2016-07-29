package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please type in your username: ")
	fmt.Scan(&name)
	fmt.Println("Your name is ", name)
}
