package main

import (
	"fmt"
	"time"
)

// worker function recieves from the jobs channel
// and send the results to the results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker:", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker:", id, "started job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// starts up to 3 workers initially blocked because there are no
	// jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// we send 5 jobs and the close the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// we collect all the results of the work.
	// THis ensures all the go routines have finished
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	// The program shows the 5 jobs being executed by our various workers
	// The program only takes 2 seconds despite doing about 5 seconds total of works
	// because there are 3 workers working concurrently

}
