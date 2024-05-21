package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func channelStop(exit <-chan struct{}) {
	for {
		select {
		case <-exit:
			fmt.Println("channelStop: DONE")
			return
		}
	}
}

func contextStop(myContext context.Context) {
	for {
		select {
		case <-myContext.Done():
			fmt.Println("contextStop: DONE")
			return
		}
	}
}

func main() {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	fmt.Println("channelStop: START")
	wg.Add(1)
	go func() {
		defer wg.Done()
		channelStop(ch)
	}()
	time.Sleep(time.Second * 2)
	ch <- struct{}{}
	wg.Wait()

	wg = sync.WaitGroup{}
	myContext, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	fmt.Println("\ncontextStop: START")
	wg.Add(1)
	go func() {
		defer wg.Done()
		contextStop(myContext)
	}()
	wg.Wait()

}
