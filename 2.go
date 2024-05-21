package main

import (
	"fmt"
	"sync"
)

// 1 способ - небуферизированный канал
func SquareNumbersChan() {
	arr := []int{2, 4, 6, 8, 10}

	intCh := make(chan int) //небуферизированный канал для int'ов. Пока не будет извлечено уже записанное значение, запись в него невозможна, и записывающая горутина блокируется
	for _, v := range arr {
		/*
			не используем исходное v, а обязательно передаём в качестве параметра (копируем в отдельную переменную), иначе т.к горутины будут обращаться к одному и тому же
			участку памяти (и запускаются не моментально), и могут быть некорретные результаты (например, цикл закончится, v=10, и все горутины будут работать с числом 10)
		*/
		go func(num int) {
			intCh <- num * num
		}(v)
	}

	//читаем данные из канала. Вывод может быть хаотичным, т.к горутины могут запускаться в работу в порядке, отличном от порядка их создания в программе
	for i := 0; i < len(arr); i++ {
		fmt.Print(<-intCh, " ")
	}
}

// 2 способ - буферизированный канал
func SquareNumbersBufferChan() {
	arr := []int{2, 4, 6, 8, 10}

	intCh := make(chan int, len(arr)) //буферизированный канал для int'ов. Пока канал полностью не заполнится на указанную ёмкость (len(arr) в данном случае), записывающая горутина блокироваться не будет
	for _, v := range arr {
		/*
			Другой вариант решения проблемы, описанной в SquareNumbersChan(). Просто создаём в каждой итерации цикла отдельную переменну vCopy, и работаем с ней в горутине.
		*/
		vCopy := v
		go func() {
			intCh <- vCopy * vCopy
		}()
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(<-intCh, " ")
	}
}

// 3 способ - WaitGroup
func SquareNumbersWg() {
	arr := []int{2, 4, 6, 8, 10}
	intCh := make(chan int, len(arr))
	var wg sync.WaitGroup
	wg.Add(len(arr)) //в группе len(arr) горутин
	for _, v := range arr {
		go func(num int) {
			defer wg.Done() //-1 к счётчику у WaitGroup
			intCh <- num * num
		}(v)
	}
	wg.Wait() //ждём выполнения всех горутин, т.е пока у WaitGroup счётчик !=0
	for i := 0; i < len(arr); i++ {
		fmt.Print(<-intCh, " ")
	}
}

// 4 способ - синхронизация вывода с помощью мьютекса и wg
func SquareNumbersMutex() {
	arr := []int{2, 4, 6, 8, 10}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range arr {
		go func(num int, mutex *sync.Mutex) {
			defer wg.Done()
			mutex.Lock()         //блокируем мютекс для вывода в консоль
			defer mutex.Unlock() //после окончания работы горутины разблокируем мьютекс
			fmt.Print(num*num, " ")
		}(v, &mutex)
	}
	wg.Wait()
}

// 5 способ - wg и канал с буфером
func SquareNumbersWgMutex() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	intCh := make(chan int, len(arr))
	wg.Add(len(arr))
	for _, v := range arr {
		go func(num int) {
			defer wg.Done()
			intCh <- num * num
		}(v)
	}
	wg.Wait()
	for i := 0; i < len(arr); i++ {
		fmt.Print(<-intCh, " ")
	}
}

func main() {
	fmt.Println("\nSquareNumbersChan")
	SquareNumbersChan()
	fmt.Println("\nSquareNumbersBufferChan")
	SquareNumbersBufferChan()
	fmt.Println("\nSquareNumbersWg")
	SquareNumbersWg()
	fmt.Println("\nSquareNumbersMutex")
	SquareNumbersMutex()
	fmt.Println("\nSquareNumbersWgMutex")
	SquareNumbersWgMutex()
}
