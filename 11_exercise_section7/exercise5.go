package main

import "fmt"

//Writing a function that can be called in all of these ways

func main(){
	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,3,2,4}
	foo(aSlice...)
	foo()
}

func foo(n ...int){
	fmt.Println(n)
}
