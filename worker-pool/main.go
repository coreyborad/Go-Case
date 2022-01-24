package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan interface{}, output chan<- interface{}) {
	for {
		select {
		case val := <-jobs:
			fmt.Println(fmt.Sprintf("job %d start on worker-%d", val, id))
			// Do job
			time.Sleep(time.Second)
			// Finish job, you can push to output chan, if main thread need catch result
			output <- val
		}
	}
}

func main() {
	jobs := make(chan interface{})
	results := make(chan interface{}, 20)
	go worker(1, jobs, results)
	go worker(2, jobs, results)

	for j := 1; j <= 20; j++ {
		fmt.Println(fmt.Sprintf("job id %d", j))
		jobs <- j
	}

	for j := 1; j <= 20; j++ {
		val := <-results
		fmt.Println(val)
	}
}
