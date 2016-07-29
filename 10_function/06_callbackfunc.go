package main

import "fmt"

func visit(numbers []int, haha func(int)) {
	for _, n := range numbers {
		haha(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers{
		if callback(n){
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	visit([]int{1, 2, 3, 4, 000000000000000000000000000}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1,2,3,4,5}, func(n int) bool{
		return n>1
	})
	fmt.Println(xs)
}

//callback is to pass in the func as the arguments

