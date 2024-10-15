package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "first"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "second"
	}()

	/*
	  We’ll use select to await both of these values simultaneously,
	  printing each one as it arrives.
	*/
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
