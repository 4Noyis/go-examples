package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
The primary mechanism for managing state in Go is communication over channels.
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/
func main() {
	/*We’ll use an atomic integer type to represent our*/
	var ops atomic.Uint64
	/*A WaitGroup will help us wait for all goroutines to finish their work.*/
	var wg sync.WaitGroup
	/*We’ll start 50 goroutines that each increment the counter exactly 1000 times.*/
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				/*To atomically increment the counter we use Add.*/
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	/*Wait until all the goroutines are done.*/
	wg.Wait()
	/*Here no goroutines are writing to ‘ops’, but using Load it’s safe to atomically read a value even while other goroutines are (atomically) updating it.*/
	fmt.Println("ops:", ops.Load())
}
