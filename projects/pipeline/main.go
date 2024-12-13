package main

import (
	"fmt"
)

// реализовать removeDuplicates(in, out chan string)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		inputStream <- "a"
		inputStream <- "a"
		inputStream <- "b"
		inputStream <- "b"
		inputStream <- "c"
		close(inputStream)
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	fmt.Print("\n")
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	prev_str := ""
	cur_str := ""
	for value := range inputStream {
		prev_str = cur_str
		cur_str = value
		if cur_str != prev_str {
			outputStream <- cur_str
		}
	}
	close(outputStream)
}
