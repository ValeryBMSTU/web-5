package main

import "sync"

func main() {
	gr := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		gr.Add(1)
		go func() {
			defer gr.Done()
			work()
		}()
	}
	gr.Wait()
}

func work() {
	// ничего
}
