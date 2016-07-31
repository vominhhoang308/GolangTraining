package main

import "fmt"

func main() {
	m := make(map[string]int) //there are three ways we can define a new map
	//var m = make(map[string]int)
	//m := map[string]int{}
	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")
	fmt.Println("map: ", m)

	v, _ := m["k1"]
	fmt.Println("and the value is ", v)
	fmt.Println(len(m))

	fmt.Println("------------------------------------------")
	// You can also declare and initialize a new map in
	// the same line with this syntax

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//these codes help you to check if the index is in the map or not, just an example ^^
	// whenever we call an index in a map, for example below, we call n["foo"], so we call the value at
	// KEY INDEX "foo", so we will always receive a VALUE at THAT INDEX and a BOOLEAN VALUE THAT INDICATES
	// whether or not that index is in range of map or not, true if it is, false if it is not.
	delete(n, "foo")

	// this code initialize (BEFORE IF statement so the scope of these two variables in only within inside the if statement)
	// VAL variable that STORES A VALUE at "foo" in the map and
	// the CHECKEXIST variable that stores BOOLEAN value that indicates if the index is inside the map or not.
	// then we will check if the index is in range, if it is then..., if it is not then..
	if val, checkExist := n["foo"]; checkExist {
		fmt.Println("val: ", val)
		fmt.Print("exists: ", checkExist)
	} else {
		fmt.Println("That value does not exist!")
		fmt.Println("That value is: ", val)
	}
}

// map is a reference type, it refers to a hash table.
// basic syntax of map, somemaps := make(map[typeOfKey]typeOfValue)
