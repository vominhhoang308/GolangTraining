package main

import "fmt"

func half(n int) (int, bool) {
	if n%2 == 0 {
		return 1, true
	} else {
		return 0, false
	}
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
