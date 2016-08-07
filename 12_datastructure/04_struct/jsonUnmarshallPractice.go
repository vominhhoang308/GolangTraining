package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 Person
	//These lines will return empty
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	//we give the slice of bytes, which is a strong of JSON.
	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)

	//we unmarshall this slice of bytes (representing Json String) back into p1, we change the value at the address p1 points to
	json.Unmarshal(bs, &p1)

	//Now p1 gets all the value from that Json string object
	fmt.Println("------------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1) //return type Person

}
