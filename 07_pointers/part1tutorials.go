package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a) //print out the address of location 43

	var b *int = &a //assign the ADDRESS to var b, b is not a type of "int pointer"

	fmt.Println(b)
	fmt.Println(*b) // * is an operator (like + - / *), it means dereference the address and give the value at that address

	*b = 42        //b says, "The value at this address, change it to 42"
	fmt.Println(a) //output 42
}

//this is useful
//we can pass a memory address instead of a bunch of values (we can pass a reference)
//and then we can stil change the value of whatever is stored at that memory address
//this makes our programs more performant
//we dont have to pass around large amounts of data
//we only have to pass around the address
//
//everythign is PASS BY VALUE in GO, like in java
//when we pass a address, we are passing a value
