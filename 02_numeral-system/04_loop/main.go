package main

import "fmt"

func main() {
	for i :=60 ; i < 161; i++ {
		//fmt.Printf("%d \t %b \t %x \n", i, i, i)
		fmt.Printf("%d \t %b \t %x \t %q \n", i , i, i, i)

	}
}
