package main

import (
	"fmt"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var m string
	for i := range inputStream {
		if m != i {
			outputStream <- i
		}
		m = i
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		inputs := []string{"apple", "banana", "apple", "orange", "banana", "grape", "grape"}
		for _, input := range inputs {
			inputStream <- input
		}
		close(inputStream)
	}()

	for unique := range outputStream {
		fmt.Println(unique)
	}
}
