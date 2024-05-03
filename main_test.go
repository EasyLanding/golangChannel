package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestHelloWorldFunc(t *testing.T) {
	var buf bytes.Buffer

	fmt.Fprint(&buf, HelloWorld())

	output := buf.String()
	expected := "Hello world!"

	if output != expected {
		t.Errorf("Unexpected output: %s", output)
	}
}

func FakeMergeChan(out chan<- int, a, b <-chan int) {
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			out <- v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out <- v
		}
	}
}

func TestMergeChan(t *testing.T) {
	// Создаем каналы
	ch1 := make(chan int)
	ch2 := make(chan int)
	mergedChan := make(chan int)

	// Ожидаемый результат
	expected := []int{1, 1, 2, 2, 3, 4, 5, 3, 4, 5}

	// Запускаем функцию mergeChan
	go func() {
		for _, v := range expected {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range expected {
			ch2 <- v
		}
		close(ch2)
	}()

	var result []int
	go func() {
		for v := range mergedChan {
			result = append(result, v)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		FakeMergeChan(mergedChan, ch1, ch2)
	}()

	wg.Wait()

	close(mergedChan)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Результат %v не равен ожидаемому %v", result, expected)
	}
}
