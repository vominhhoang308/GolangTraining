package main

import (
	"fmt"
	"math"
	"strconv"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 { return math.Pi * c.radius * c.radius } // this function takes A VALUE AS A RECEIVER
//func (c *circle) area() float64{ return math.Pi * c.radius * c.radius } // this function takes A POINTER AS A RECEIVER
//func (s square) area() float64{ return s.side*s.side }  // this function takes A VALUE AS A RECEIVER
func (s *square) area() float64 { return s.side * s.side } // this function takes A POINTER AS A RECEIVER

func info(s shape) string {
	return strconv.FormatFloat(s.area(), 'f', 6 , 64)
}

func main() {
	c := circle{5}
	s := square{5}

	// These two functions below will work as we declared in the function above that we take the VALUE as the RECEIVER
	// and no matter what the parameter is, the function will always work. So it can be either VALUE or POINTER as the parameters
	// if this is a VALUE, then GOLANG will compile it straight away
	// if this is a POINTER (it is *c in this case) then GOLANG compiler will automatically deference it and gave it a value
	fmt.Println("This is the function that takes the VALUE as the parameters for info() for circle: ", info(c))
	fmt.Println("This is the function that takes the RECEIVER as the parameters for info() for circle: ", info(&c))

	// The first statement does not work but the second statement works.
	// As we declare in the function area() above, we takes only a POINTER RECEIVER.
	// If we parse in a POINTER PARAMETERS (like in second statement), GOLANG compiler will understand and run it
	// However, when we parse in a VALUE PARAMTERS for a POINTER RECEIVER, then GOLANG will gives out errors.
	// The reason behind it is that we give GOLANG a value, GOLANG cannot find the adderss of it.
	// For example in real life situation, we give ourselves a PENTHOUSE, then we cant never find the address of that PENTHOUSE,
	// but if we give ourselves the address, then we can find what type of house is that address store. (this is how we can give
	// the POINTER and GOLANG can deference and the value for us).

	//fmt.Println("This is the function that takes the VALUE as the parameters for info() for square: ", info(s))
	fmt.Println("This is the function that takes the RECEIVER as the parameters for info() for square: ", info(&s))
}
