package main

import "fmt"

// реализовать removeDuplicates(in, out chan string)
func removeDuplicates(inputStream chan string, outputStream chan string) {
	str := ""
	for v := range inputStream {
		if str == "" {
			str = v
			outputStream <- v
		} else if str != v {
			str = v
			outputStream <- v
		} else if str == v {
			continue
		}
	}
	close(outputStream)
}

func main() {
	// здесь должен быть код для проверки правильности работы функции removeDuplicates(in, out chan string)
	inputStream := make(chan string)
	outputStream := make(chan string)

	go func() {
		inputStream <- "p"
		inputStream <- "b"
		inputStream <- "b"
		inputStream <- "r"
		close(inputStream)
	}()

	go removeDuplicates(inputStream, outputStream)

	for value := range outputStream {
		fmt.Println(value)
	}
}
