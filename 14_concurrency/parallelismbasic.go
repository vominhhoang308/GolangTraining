package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2) // this is to add 2 threads in a waitgroup
	go foo()  // this is to add a go routine, aka. threads
	go bar()
	wg.Wait() // this will wait until a waitgroup becomes 0
}

func foo() {
	for i := 0; i < 40; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done() // when this is done, it will minus 1 from the waitgroup, so it basically means that it will stop the thread if the thread finish running
}

func bar() {
	for i := 0; i < 40; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done() // same thing
}
