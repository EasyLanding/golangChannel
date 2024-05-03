package main

import (
	"fmt"
	"sync"
	"time"
)

func HelloWorld() string {
	return "Hello world!"
}

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(from))

	for _, ch := range from {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				mergeTo <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()
}

func mergeChan2(chans ...chan int) chan int {
	mergedChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, ch := range chans {
		go func(ch chan int) {
			defer wg.Done()
			for val := range ch {
				mergedChan <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func generateData(n int) chan int {
	dataChan := make(chan int)

	go func() {
		defer close(dataChan)
		for i := 0; i < n; i++ {
			dataChan <- i
		}
	}()

	return dataChan
}

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func timeout(timeout time.Duration) func() bool {
	ch := make(chan struct{})

	go func() {
		time.Sleep(timeout)
		close(ch)
	}()

	return func() bool {
		select {
		case <-ch:
			return false
		default:
			return true
		}
	}
}

func main() {
	ch1 := generateChan(3)
	ch2 := generateChan(4)
	ch3 := generateChan(2)

	mergedChan := make(chan int)
	go mergeChan(mergedChan, ch1, ch2, ch3)

	for val := range mergedChan {
		fmt.Println(val)
	}

	// // Для mergeChan2
	mergedChan2 := mergeChan2(ch1, ch2, ch3)
	for val := range mergedChan2 {
		fmt.Println(val)
	}

	data := generateData(10)
	go func() {
		time.Sleep(1 * time.Second) // ждем одну секунду, чтобы горутина main успела выполниться
		close(data)
	}()
	for num := range data {
		fmt.Println(num)
	}

	ch := make(chan int, 1)

	// Попытка отправить значение в канал
	fmt.Println("Попытка отправить значение 42 в канал...")
	if trySend(ch, 42) {
		fmt.Println("Значение успешно отправлено в канал.")
	} else {
		fmt.Println("Не удалось отправить значение в канал, так как он заполнен.")
	}

	// Попытка отправить другое значение в канал
	fmt.Println("Попытка отправить значение 100 в канал...")
	if trySend(ch, 100) {
		fmt.Println("Значение успешно отправлено в канал.")
	} else {
		fmt.Println("Не удалось отправить значение в канал, так как он заполнен.")
	}

	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
