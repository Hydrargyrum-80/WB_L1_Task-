package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int //непосредственно счётчик
	mutex sync.Mutex
}

func (c *counter) init() {
	c.count = 0
	c.mutex = sync.Mutex{}
}

func (c *counter) increment() {
	c.mutex.Lock() //блокируем доступ к счётчику. Теперь операция инкремента доступна только в одном месте.
	//остальные вызовы - ждут
	defer c.mutex.Unlock()
	c.count++
}

func main() {
	wg := sync.WaitGroup{}
	count := counter{}
	count.init()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				count.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.count)
}
