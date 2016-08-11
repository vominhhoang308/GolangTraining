package main

import (
	"errors"
	"log"
)

func Sqrt(f float64) (float64, error) { // here we create our new function called Sqrt and we return the float64, error type
	// error in here is the interface.
	if f < 0 {
		// package errors have the function New(s string) (errors) and return the error interface (notice the differences
		// between package errors and error interface, with s and without s)
		// errors.New actually return a pointer to a error string and assign that to our error in our function
		return 0, errors.New("incorrect input: square root of negative number")
	}
	return 100, nil // so if the condition is correct, we return the answer (answer here is just dummy text) and nil error message

}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}
