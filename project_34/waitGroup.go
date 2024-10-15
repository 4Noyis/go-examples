/*
To wait for multiple goroutines to finish, we can use a wait group.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)

	fmt.Printf("worker %d done\n", id)
}

func main() {
	/*This WaitGroup is used to wait for all the goroutines launched here to finish.*/
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		/*Launch several goroutines and increment the WaitGroup counter for each.*/
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	/*Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.*/
	wg.Wait()
}
