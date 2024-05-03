package main

import (
	"fmt"
	"sync"
)

func HelloWorld() string {
	return "Hello world!"
}

func mergeChan(mergeTo chan int, from ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(from))

	for _, ch := range from {
		go func(ch chan int) {
			defer wg.Done()
			for val := range ch {
				mergeTo <- val
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

func main() {
	ch1 := generateChan(3)
	ch2 := generateChan(4)
	ch3 := generateChan(2)

	mergedChan := make(chan int)
	go mergeChan(mergedChan, ch1, ch2, ch3)

	for val := range mergedChan {
		fmt.Println(val)
	}

	// Для mergeChan2
	mergedChan2 := mergeChan2(ch1, ch2, ch3)
	for val := range mergedChan2 {
		fmt.Println(val)
	}
}
