package main

import "fmt"

func greatest(number ...int) int {
	var greatestNumber int
	for _, v := range number {
		if v > greatestNumber {
			greatestNumber = v
		}
	}
	return greatestNumber
}

func main() {
	i := greatest(1, 2, 5000, 4, 5, 6, 7)
	fmt.Println(i)
}
