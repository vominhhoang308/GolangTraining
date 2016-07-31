package main

import "fmt"

// We create a new TYPE named Person (new class in other languages)
type Person struct {
	First string
	Last  string
	Age   int
}

// We create a new type named DoubleZero that is Reusability (embedded) from type Person, another word
// Person is embedded into Double Zero
// Other languages, we would say DoubleZero is inherited from Person, it has all the methods and the states
// In go we can even use Polymorphism, that means we can create a new "fields" (other languages: states) that can OVERWRITE the inner struct (Person)
// In this case, we can overwrite the First and Last fields in the inner struct Person
// However, we can still access the First and Last fields OF THE INNER STRUCT PERSON, example at the end of the code.

type DoubleZero struct {
	Person
	First         string
	Last          string
	LicenseToKill bool
}

// We create a greeting function for Person
func (p Person) Greetings() string {
	return "Hello " + p.First + p.Last
}

// We overwriting the Greeting function for Dz.
func (dz DoubleZero) Greetings() string {
	return "Hello " + dz.First + dz.Last
}

func main() {
	// We create a variables p1 with type Person (we crate an instant p1 with class Person in other languages)
	var p1 Person = Person{
		First: "James",
		Last:  "Bond",
		Age:   29,
	}

	var p2 DoubleZero = DoubleZero{
		Person: Person{
			First: "Minh Hoang",
			Last:  "Vo",
			Age:   20,
		},
		First:         "Minh Hoang 00",
		Last:          "Vo 00",
		LicenseToKill: true,
	}

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill)
	fmt.Println(p2.Person.First, p2.Person.Last, p2.Person.Age) // we have access to First and Last of TYPE PERSON inside p2 which is type DoubleZero
	// In Java, we can say that we create an instant of DoubleZero that is REFER to Person
	// NB: notice the Age fields, we can access to Age of p2 by both DoubleZero(p2.Age) or Person (p2.Person.Age)
	fmt.Println("---------------------------------")
	/////////////////////////////////

	fmt.Println(p1.Greetings())        // We call greeting function for a Person type p1.
	fmt.Println(p2.Greetings())        // We call greeting function for a DZ type p2, this will call a GREETING THAT IS OVERWRITED FOR DOUBLE ZERRO TYPE
	fmt.Println(p2.Person.Greetings()) // We call greeting function for DZ type P2, HOWEVER, this will call a GREETING FUNCTION IN PERSON TYPE, so it will print out is inside Person of DZ
}
