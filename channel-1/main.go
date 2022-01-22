package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	doneChan := make(chan struct{})
	countChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(val int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println("finished job id:", val)
			countChan <- val
		}(i, &wg)
	}
	go func() {
		wg.Wait()
		close(doneChan)
	}()
	count := 0
	for {
		select {
		case <-doneChan:
			return
		case val := <-countChan:
			fmt.Println("finished:", val)
			count++
			fmt.Println(count)
		}
	}
}
