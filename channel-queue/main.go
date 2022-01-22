package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(jobQueue <-chan int) {
		for {
			select {
			case val := <-jobQueue:
				fmt.Println("finished:", val)
				time.Sleep(time.Second)
			}
		}
	}(queue)

	fmt.Println(send(1, queue), 1)
	time.Sleep(time.Second)
	fmt.Println(send(2, queue), 2)
	fmt.Println(send(3, queue), 3)
	fmt.Println(send(4, queue), 4)
	time.Sleep(time.Second)
	fmt.Println(send(5, queue), 5)
	time.Sleep(time.Second * 3)
}

func send(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
