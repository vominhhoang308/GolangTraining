package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // So the code STOP RIGHT HERE UNTIL THERE IS SOMETHING THAT TAKES THE VALUE OFF THE CHANNEL
		}
		close(c) // this function closes the channel, so no other value will be put in after this
	}()

	for n := range c {
		fmt.Println(n)
	}

	// we do not need this: "time.Sleep(time.Second)" anymore because we closed it down and we ran it in main routines.
}
