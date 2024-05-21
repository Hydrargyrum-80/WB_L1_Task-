package main

import (
	"fmt"
	"sync"
	"time"
)

func customSleep(duration time.Duration) {
	p1 := time.Now().Add(duration) //фиксируем текущее время + duration
	for {
		if time.Now().After(p1) { //ждём пока время не превысит p1
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(1)
	//проверяем точность
	go func() {
		defer wg.Done()
		customSleep(time.Second * 5)
	}()
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Elapsed time: ", duration)
}
