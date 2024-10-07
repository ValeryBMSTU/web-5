package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	a := make(chan int)
	go func(b chan int) {
		defer close(b)
		select {
		case c := <-firstChan:
			b <- c * c
		case d := <-secondChan:
			b <- d * 3
		case <-stopChan:
			return
		}

	}(a)

	return a
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	r := calculator(ch1, ch2, stop)
	//ch1 <- 3
	//ch2 <- 4
	close(stop)
	fmt.Println(<-r)
}
