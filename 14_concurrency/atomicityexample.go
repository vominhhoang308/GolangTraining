package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// ATOMICITY is sort of MUTEX but instead of locking a whole
// chunk of codes
// it locks only one variable.

var wg sync.WaitGroup
var counter int64

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
		atomic.AddInt64(&counter, 1) // this will lock value COUNTER, so one thread can access to it.
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
