package main

// MUTEX stands for mutually exclusion object
// when we use MUTEX in this case, var COUNTER can only be access by one thread
// at the time, it is still a race condition,
// but it will eventually not be overwritten.

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		mutex.Lock() // this will lock all the var below, so only 1 thread can access to it
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock() // this will unlock the object to let other thread access to it
	}
	wg.Done()
}
