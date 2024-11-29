package main

import (
	"fmt"
	"sync"
)

func work() {
	// Здесь ваш код, который выполняется в каждой горутине
	fmt.Println("WORK!")
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup для синхронизации

	wg.Add(10) // Увеличиваем счетчик на 10, так как запускаем 10 горутин

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик после завершения горутины
			work()          // Вызываем функцию work
		}()
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("Все горутины завершились!")
}
