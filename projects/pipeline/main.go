package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream) // Закрываем выходной канал при завершении

	var prevValue string // Храним предыдущее значение
	for value := range inputStream {
		if value != prevValue { // Проверка на дубликат
			outputStream <- value // Отправка в выходной канал, если значение не дубликат
			prevValue = value     // Обновляем предыдущее значение
		}
	}
}

func main() {
	in := make(chan string)
	out := make(chan string)

	go func() {
		defer close(in) // Закрываем входной канал при завершении
		in <- "one"
		in <- "two"
		in <- "two"
		in <- "three"
		in <- "three"
		in <- "three"
		in <- "four"
		in <- "five"
	}()

	go removeDuplicates(in, out)

	for value := range out {
		fmt.Println(value)
	}
}
