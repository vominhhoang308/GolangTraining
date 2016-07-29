package main

import "fmt"

func world() {
	fmt.Println("world!")
}

func hello() {
	fmt.Print("Hello ")
}

func main() {
	defer world()
	hello()
}

// DEFER in line 14 basically means POSTPONE, so wherever we type defer before the arguments or the statement, it would
//postpone right until the surrounding outer function exit.
//in this case, the program will pass on to run hello(), and right before main() exit, it runs world()
