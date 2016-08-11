package main

import (
	"errors"
	"log"
)

// Just like how we create our custom err message inside function Sqrt() in file customFucntionWithErr.go
// now we put the err message outside of function and assign that to a larger scope so we can use it anywhere we want
// we say it IdealWay, but it is actually a must for better programming
// Naming convention should be like "Err....", start with "Err" and then something else.
// remember the capital "E" so we can use the Err anywhere in the projects
// this is more Abstract, and anything with Abstraction is better programming.

var ErrIncorrectSqrt = errors.New("Sqrt(): incorrect input, cannot squareroot a negative")

func SquareRoot(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrIncorrectSqrt
	}
	return 42, nil
}

func main() {
	_, err := SquareRoot(-10)
	if err != nil {
		log.Fatal(err)
	}

}

// Couple of examples of how IDEALLY it is (using errors.New() in standard library)
// http://golang.org/src/pkg/buffio/buffio.go
// http://golang.org/src/pkg/io/io.go
