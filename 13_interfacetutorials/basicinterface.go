package main

import (
	"fmt"
	"math"
)

// We are creating a Square type struct (in other language, Square Class)
type Square struct {
	side float64
}

// We are creating another Circle type struct
type Circle struct{
	radius float64
}

// We are creating a Shape with type interface with the method area(), all the struct type that
// Implement this interface will then have to implement this area() method
// Unlike in java, where you would have something like: public Square implements Shape{}
// In GO, the only thing to show that you are implementing this interface is too implement the method of this interface with difference receiver
type Shape interface {
	area() float64
}

// This method takes a Shape interface as the parameters, that means it can take any type that is implements Shape
func info(z Shape) {
	fmt.Println(z)

	fmt.Println("Area of the shapes ", z.area())
}

// We are implementing Shape interface to a Square struct by implementing method area()
func (s Square) area() float64 {
	return s.side * s.side
}

// We are implementing Shape interface to a Circle struct by implementing method area()
func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func main() {
	var squareExample Square = Square{15}

	fmt.Println(squareExample.area())
	fmt.Println("-------------------------------")
	info(squareExample)

	fmt.Println("-------------------------------")

	var circleExample Circle = Circle{15}
	info(circleExample)
}
