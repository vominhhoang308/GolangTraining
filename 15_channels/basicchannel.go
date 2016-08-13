package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // We create a channel using MAKE, because it is a reference type, like map and slice
				// this is an UNBUFFERED CHANNEL, which means that there would be no limitation in terms of space

	// explanation of this code below: we have two thread running concurrency, we call it go routine running concurrency
	// we already created a channel
	// we loop through 0 to 10
	// however, at each value, we pass that value into a channel
	// then THAT LOOP WILL NOT CONTINUE IF THAT VALUE DOES NOT GOES TO ANOTHER THREAD
	// so it basically means that,
	// when i=0, it pass to c, then we have to print c, then i=1 will pass to c
	// loop will wait for the value to run, and the channel print whenever the new value is passed in
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second);
}
