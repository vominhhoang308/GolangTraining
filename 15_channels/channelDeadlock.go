package main

import ()
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	//getData(ch) // this will cause the deadlock error. Go Runtime found out that we have one go routines that need some actions to it.
	time.Sleep(1e9)

}

func sendData(ch chan string) {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
}

func getData(ch chan string) {

	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
