package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := &Person{"JamesBond", 22} // we set p1 TO A POINTER TO A NEWLY CREATED VARIABLES TYPE PERSON{JamesBond, 22}
	fmt.Println(p1) //output is the address to somewhere that stores {jamesbond, 22}, actual output: &{"JamesBond",22}, we call it struct pointer
	fmt.Printf("%T\n", p1) // We check the type of p1, output: *main.Person, so it is at type Struct Pointer, or to be exact, it is at type main.Person Pointer,
				// things with the same type main.Person Pointer (e.g. another variables var1)would definitely have type(*var1) = main.Person

	fmt.Printf("%T \n", *p1) // this is to check the type of (*p1), output type main.Person

	fmt.Println(p1.name) // this will give you JamesBond, so GO compiler auto convert your code into fmt.Println(*p1.name), so it will give you the value at the address p1.name
	fmt.Println(p1.age)
}
