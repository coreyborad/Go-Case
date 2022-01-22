package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(val int) {
			fmt.Println(val)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
