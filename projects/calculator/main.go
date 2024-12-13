package main

import "fmt"

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int

func main() {
	chan1, chan2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	r := calculator(chan1, chan2, stop)
	go func() {
		chan1 <- 3
		chan2 <- 3
		close(stop)
	}()
	fmt.Println(<-r)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	res := make(chan int)
	go func() {

		select {
		case num := <-firstChan:
			res <- num * num
		case num := <-secondChan:
			res <- num * 3
		case _ = <-stopChan:
		}
		close(res)

	}()
	return res
}
