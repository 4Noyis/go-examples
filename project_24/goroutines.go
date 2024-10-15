package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread of execution.
*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
	/*
		  When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
		The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
	*/
}
