package main

import (
	"fmt"
	"sync"
)

func main() {
	dataIntChan := make(chan int)
	resultIntChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(outChan chan<- int) {
		defer wg.Done()
		defer close(outChan)
		var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, i := range numbers {
			outChan <- i
		}
	}(dataIntChan)

	go func(inChan <-chan int, outChan chan<- int) {
		defer wg.Done()
		defer close(outChan)
		for i := range inChan {
			outChan <- i * i
		}
	}(dataIntChan, resultIntChan)
	for i := range resultIntChan {
		fmt.Println("Result:", i)
	}
	wg.Wait()
}
