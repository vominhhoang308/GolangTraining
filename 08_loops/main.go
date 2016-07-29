package main

import "fmt"

func main() {
	var i int = 0
	for  {
		i++
		if i%2==0{
			continue
		}
		fmt.Println(i)
		if i>=50{
			break
		}
	}
}
