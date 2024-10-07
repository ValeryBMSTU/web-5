package main

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var a string
	for i := range inputStream {
		if i != a {
			outputStream <- i
			a = i
		}
	}
	close(outputStream)
}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go func() {
		inputStream <- "Hi"
		inputStream <- "Alena"
		inputStream <- "Hi"
		inputStream <- "Hi"
		inputStream <- "Hi"
		inputStream <- "!"
		inputStream <- "!"
		close(inputStream)
	}()
	go removeDuplicates(inputStream, outputStream)

	for v := range outputStream {
		fmt.Print(v)
	}
}
