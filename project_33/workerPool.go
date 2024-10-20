package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)

		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	/*This starts up 3 workers, initially blocked because there are no jobs yet.*/
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	/*Here we send 5 jobs and then close that channel to indicate that’s all the work we have.*/
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	/*
	   Our running program shows the 5 jobs being executed by various workers.
	   The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.
	*/
}
