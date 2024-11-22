package main

import (
	"fmt"
	"time"
)

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	var prevValue string
	firstValue := true

	for val := range inputStream {
		if firstValue || val != prevValue {
			outputStream <- val
			prevValue = val
			firstValue = false
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)
		inputStream <- "apple"
		inputStream <- "apple"
		inputStream <- "banana"
		inputStream <- "banana"
		inputStream <- "banana"
		inputStream <- "cherry"
		inputStream <- "apple"
	}()

	go func() {
		for result := range outputStream {
			fmt.Println(result)
		}
	}()

	time.Sleep(1 * time.Second)
}
