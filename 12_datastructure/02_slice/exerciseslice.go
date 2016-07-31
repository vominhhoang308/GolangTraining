package main

import "fmt"

func main() {
	records := make([]string, 5, 10)
	appendedslice := ([]string{"abc", "def"})
	fmt.Println(len(records))
	fmt.Println(cap(records))

	fmt.Println("------------------------------")
	records = append(records, appendedslice...)

	fmt.Println(len(records))
	fmt.Println(records)

	fmt.Println("------------------------------")
	records = append(records[:5], records[6:]...)
	fmt.Println(len(records))
	fmt.Println(records)
}
