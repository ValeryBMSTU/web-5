package main

import (
	"fmt"
	"time"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)
		select {
		case val := <-firstChan:
			resultChan <- val * val
		case val := <-secondChan:
			resultChan <- val * 3
		case <-stopChan:
			return
		}
	}()

	return resultChan
}

func main() {

	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		firstChan <- 4
	}()

	result := <-calculator(firstChan, secondChan, stopChan)
	fmt.Printf("Тест 1: Получено %d, ожидается 16\n", result)

	go func() {
		secondChan <- 5
	}()

	result = <-calculator(firstChan, secondChan, stopChan)
	fmt.Printf("Тест 2: Получено %d, ожидается 15\n", result)

	go func() {
		stopChan <- struct{}{}
	}()

	resultChan := calculator(firstChan, secondChan, stopChan)
	select {
	case result = <-resultChan:
		fmt.Printf("Тест 3: Неожиданный результат %d\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Тест 3: Канал корректно закрылся")
	}
}
