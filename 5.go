package main

import (
	"fmt"
	"sync"
	"time"
)

func timeStopChannel(N time.Duration) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 2)
	go func(ch chan<- int) {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}(ch)
	go func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Read: ", v)
		}
	}(ch)
	time.Sleep(N * time.Second)
	wg.Wait()
}

func main() {
	timeStopChannel(5)
}
