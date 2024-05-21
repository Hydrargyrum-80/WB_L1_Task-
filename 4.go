package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(workerId int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch { //worker функционирует пока открыт канал для получения данных
		fmt.Println("Worker ", workerId, " - ", v)
	}
	fmt.Println("Worker ", workerId, " - DONE")
}

func main() {
	workersCount := 5
	intCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ { //Запускаем нужное кол-во worker'ов
		go worker(i, intCh, &wg)
	}
	go func() { //Основная горутина, записывающая данные в канал
		var chanVal int
		for i := 0; i < 256; i++ {
			chanVal = rand.Int()
			intCh <- chanVal
			time.Sleep(1 * time.Millisecond)
		}
	}()
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	//До ввода Ctrl+C основная горутина заблокирована
	<-exitSignal
	close(intCh) //Закрываем канал для отправки данных
	/*
		Дожидаемся корректного завершения работы всех worker'ов. Сразу завершать работу нельзя, т.к в не завершённых горутинах
		могут остаться необработанные/несохраненные данные/файлы, незакрытые соединения, запросы и т.д, и принудительная остановка
		программы приведёт к их потере и ошибкам
	*/
	wg.Wait()
}
