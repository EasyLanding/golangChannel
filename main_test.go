package main

import (
	"bytes"
	"fmt"
	"reflect"
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

func TestMergeChan(t *testing.T) {
	ch1 := generateChan(3)
	ch2 := generateChan(4)
	ch3 := generateChan(2)

	mergedChan := make(chan int)
	go mergeChan(mergedChan, ch1, ch2, ch3)

	results := make([]int, 0)
	for val := range mergedChan {
		results = append(results, val)
	}

	expected := results
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", results, expected)
	}
}

func TestMergeChan2(t *testing.T) {
	ch1 := generateChan(3)
	ch2 := generateChan(4)
	ch3 := generateChan(2)

	mergedChan := mergeChan2(ch1, ch2, ch3)

	results := make([]int, 0)
	for val := range mergedChan {
		results = append(results, val)
	}

	expected := results
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Unexpected result. Got: %v, Expected: %v", results, expected)
	}
}
