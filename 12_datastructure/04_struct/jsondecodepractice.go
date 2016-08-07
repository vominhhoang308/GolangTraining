package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	// decoding is the process of getting some JSON stream in the outer source
	// (in web, getting api somewhere that returns a json string object)
	// and then decode that JSON string into our readable object that we can access to.

	var p1 Person // we creating an EMPTY p1 with type Person

	// now we are getting an JSON object and decode it, put that into p1
	// (we are actually put that into a value that has the address that p1 refers to, &p1)

	// our json object is (`"First":"James", "Last":"Bond", "Age":22`)
	// in order to decode it, we have to use Decode which will take the parameters of type DECODE POINTER
	// in order to have that type DECODER POINTER, we need to call a newDecoder
	// newDecoder takes READER INTERFACE as the parameters, so we need a READER INTERFACE
	// in order to have an reader interface, we need to CHANGE OUR JSON STRING INTO READER, so..
	// we need to create a variable using NewReader to turn that JSON STRING TO READER variables
	// NewReader actually takes a string as the parameters and return a READER POINTER
	// READER POINTER implements Reader interface, so we HAVE READER INTERFACE
	// now we can use newDecoder on it, and then Decode it

	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":22}`) //rdr stands for reader
	var decodeObject = json.NewDecoder(rdr)                                //we make a DECODE pointer and assign that to decodeObject
	decodeObject.Decode(&p1)                                               //we decode that decodeObject and put that value into the value with the address of P1

	//Now we have p1 : Person{"James", "Bond", "22"}
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
