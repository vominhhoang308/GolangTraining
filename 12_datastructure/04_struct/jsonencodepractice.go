package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	//creating a variable p1 for type main.Person, then we will encode that into Json object and send it to os.Stdout
	p1 := Person{"James", "Bond", 22}

	//these two functions creates a new ENCODER POINTER, then use that to encode p1 into json string object

	//NewEncoder takes a WRITER INTERFACE as the params, os.Stdout implements WRITER INTERFACE so it is usable in this concept
	//NewEncoder returns the ENCODE POINTER that WRITE to the value in params, in this case, os.Stdout
	//(os.Stdout) is just a place to write or read data temporary
	var encodeObject = json.NewEncoder(os.Stdout)

	//Once we have the varibles which is of type ENCODE POINTER, then we can use function Encode with the parammeters is an interface
	//This will take value in parameters, write it in a JSON format and parse it to the stream, return error if theres any error.
	encodeObject.Encode(p1)
}
