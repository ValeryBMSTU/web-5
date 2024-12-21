package main

import "fmt"

// реализовать calculator(firstChan <-chan int, secondChan <-chan int,
// stopChan <-chan struct{}) <-chan int
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	res := make(chan int)
	var v int
	go func() {
		defer close(res)
		select {
		case v = <-firstChan:
			res <- v * v
		case v = <-secondChan:
			res <- v * 3
		case <-stopChan:
			return
		}
	}()
	return res
}
func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	r := calculator(ch1, ch2, stop)
	//ch1 <- 4
	ch2 <- 3
	close(stop)
	fmt.Println(<-r)
}
