package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	outputChan := make(chan int)

	go func() {
		defer close(outputChan)
		select {
		case val, ok := <-firstChan:
			if ok {
				outputChan <- val * val
			}
		case val, ok := <-secondChan:
			if ok {
				outputChan <- val * 3
			}
		case <-stopChan:
			return
		}
	}()
	return outputChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		stopChan <- struct{}{}
	}()
	outputChan := calculator(firstChan, secondChan, stopChan)

	for result := range outputChan {
		fmt.Println("Результат:", result)
	}
}
