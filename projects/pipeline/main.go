package main

import (
	"fmt"
	"time"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	o := ""
	for i := range inputStream {
		if o != i {
			outputStream <- i
			o = i
		} else if o == i {
			continue
		}
	}
	close(outputStream)
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go printer(outputStream)
	for _, i := range "112334456" {
		inputStream <- string(i)
	}
	var input string
	fmt.Scanln(&input)
}
