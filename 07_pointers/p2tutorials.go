package main

import "fmt"

// this is the wrong version to show how GO is PASS BY VALUE and how to use POINTERS in GO.

func zero(z int) {
	z = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // OUTPUT still 5
}

// The reason behind in details is that
// in main, when we call the function with the parameters x, it does not passing the reference of x
// what it does is it only passes the value 5, so it would rather be called zero(5)
// then it goes to called zero(5), so that function assign it own z to value 5, changing z to value 0.
// it has nothing to do with the original x in the main

// this is the correct way to implement this function by using pointer.

func zero(z *int){
	*z = 0
}

func main(){
	x := 5
	zero(&x)
	fmt.Println(x) // OUTPUT 0
}

//function zero takes a INT POINTER TYPE as the parameters, then it changed the value that has that MEMORY ALLOCATION ADDRESS
//into 0 (that address can store any kind of number type int)
//in main, x is assign to have value 5, then ITS ADDRESS (address of value 5) is passed in as parameters
//then the function changing the value, that has the ADDRESS passed in, to 0.
//Now the address of value of x store only 0, so x returns 0.
