package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("--------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("---------------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), " Cappacity ", cap(mySlice), " Value: ", mySlice[i])
	}

}

// It prints out the empty slice at first, with length of 0 and the capacity is 5.

// After the loop, the length increased by 1 after each loop, once it reaches the capacity, the capacity actually increased by doubles each time until it can store all the elements.
