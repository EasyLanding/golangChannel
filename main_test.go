package main

import (
	"bytes"
	"fmt"
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

func TestGenerateData(t *testing.T) {
	data := generateData(5)

	count := 0
	for num := range data {
		if num != count {
			t.Errorf("Expected %d, but got %d", count, num)
		}
		count++
	}

	if count != 5 {
		t.Errorf("Expected 5 numbers, but got %d", count)
	}
}
