package main

import (
	"fmt"
	"sync"
)

func SquareNumbersChanSum() {
	arr := []int{2, 4, 6, 8, 10}
	intCh := make(chan int)
	for _, v := range arr {
		go func(num int) {
			intCh <- num * num
		}(v)
	}
	squareSum := 0
	for i := 0; i < len(arr); i++ {
		squareSum += <-intCh
	}
	fmt.Println(squareSum)
}

func SquareNumbersBufferChanSum() {
	arr := []int{2, 4, 6, 8, 10}
	intCh := make(chan int, len(arr))
	for _, v := range arr {
		go func(num int) {
			intCh <- num * num
		}(v)
	}
	squareSum := 0
	for i := 0; i < len(arr); i++ {
		squareSum += <-intCh
	}
	fmt.Println(squareSum)
}

func SquareNumbersWgMutexSum() {
	arr := []int{2, 4, 6, 8, 10}
	mutex := sync.Mutex{}
	squareSum := 0
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go func(num int, mutex *sync.Mutex) {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
			squareSum += num * num
		}(v, &mutex)
	}
	wg.Wait()
	fmt.Println(squareSum)
}

func main() {
	fmt.Println("SquareNumbersChanSum")
	SquareNumbersChanSum()
	fmt.Println("SquareNumbersBufferChanSum")
	SquareNumbersBufferChanSum()
	fmt.Println("SquareNumbersWgMutexSum")
	SquareNumbersWgMutexSum()
}
