package main

import "fmt"

// calculator - функция, которая принимает три канала:
// firstChan, secondChan, stopChan, и возвращает канал resChan типа <-chan int.
// Она запускает горутину, которая обрабатывает значения из каналов.
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resChan := make(chan int) // Создаем канал для результата

	go func() { // Запускаем анонимную горутину
		defer close(resChan) // Закрываем канал при завершении

		// select позволяет ожидать значения из любого из каналов
		select {
		// Получение значения из firstChan
		case val := <-firstChan: // Получаем значение из канала
			resChan <- val * val // Квадрат значения и отправка в resChan
		// Получение значения из secondChan
		case val := <-secondChan: // Получаем значение из канала
			resChan <- val * 3 // Умножение значения на 3 и отправка в resChan
		// Получение сигнала из stopChan
		case <-stopChan: // Получаем сигнал из канала
			return // Завершаем горутину
		}
	}()

	return resChan // Возвращаем канал результата
}

func main() {
	ch1 := make(chan int)      // Создаем канал ch1
	ch2 := make(chan int)      // Создаем канал ch2
	ch3 := make(chan struct{}) // Создаем канал ch3 (для сигнала остановки)

	res := calculator(ch1, ch2, ch3) // Вызываем функцию calculator и получаем канал res

	ch2 <- 16 // Отправляем значение 16 в ch2

	fmt.Println(<-res) // Получаем результат из канала res и выводим его
}
