package main

import (
	"fmt"
	"time"
)

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ch := make(chan int)
	var val int

	go func() {
		defer close(ch)
		select {
		case val = <-firstChan:
			ch <- val * val
		case val = <-secondChan:
			ch <- val * 3
		case <-stopChan:
			return
		}
	}()
	return ch
}
func main() {

	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	outputChan := calculator(firstChan, secondChan, stopChan)

	go func() {
		time.Sleep(5 * time.Second)
		close(stopChan)
	}()

	go func() {
		//firstChan <- 5
		secondChan <- 10
		firstChan <- 21
		secondChan <- 13
	}()

	for result := range outputChan {
		fmt.Println(result)
	}
}
