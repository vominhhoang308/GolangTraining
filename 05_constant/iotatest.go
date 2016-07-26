package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b\t",MB)
	fmt.Printf("%d\n",MB)
	fmt.Printf("%b\t",KB)
	fmt.Printf("%d\n",KB)
	fmt.Printf("%d\n",TB)
	fmt.Printf("%b\t",TB)
}
