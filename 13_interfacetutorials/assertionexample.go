package main

import "fmt"

// CONVERSION WORKS WITH STRING FLOAT64 INT32 UINT8 etc., pretty much primitives type, for string, we have another package name strconv

// ASSERTION WORKS WITH INTERFACE, we convert AN INTERFACE INTO int, string,..etc
func main() {
	rem := 7.24
	fmt.Printf("%T \n", rem)      // this is should be of TYPE float64
	fmt.Printf("%T \n", int(rem)) // we CONVERT type float64 into AN TYPE INT

	// this is for assertion
	var val interface{} = 7
	fmt.Printf("%T \n", val) // this is of type int, eventho it is assign to an EMPTY INTERFACE,
				// however, we still cannot do this: val + 6, it will give out the MISMATCH TYPE errors.
	//fmt.Println(val + 6) // returns error

	//fmt.Printf("%T \n", int(val)) // this will give out the error, we cannot CONVERT AN INTERFACE INTO STRING
					// quick NOTICE WITH THE SYNTAX, convertion is to with the syntax int(some_var) to convert some other primitive type to int
	fmt.Printf("%T \n", val.(int)) // this will  give out type INT, notice the SYNTAX,
					// we put the some_var the left hand and the type we want to convert into on the right hand
					// in the bracket
					// and we can do       val.(int) + 6           it will give out 13.
	fmt.Println(val.(int) + 6) // returns 13
}
