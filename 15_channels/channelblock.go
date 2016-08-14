package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	//fmt.Println(<-ch1) // prints only 0

	go pumpReceiver(ch1)

	time.Sleep(1e9)
}

// Because there is no receiver, so only 0, which is the first value that parse in the channel, will be printed.

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

// in order to fix it, we need to make this receiver function.

func pumpReceiver(ch chan int) {
	for {
		//varReceiver := <- ch
		fmt.Println(<-ch)
	}
}
