package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Order - заказ.
type Order struct {
	ID int
}

// функция получения и обработки заказа., подсчет количества проведенных заказов
func Worker(id int, orders <-chan Order, wg *sync.WaitGroup, counters []int) {
	defer wg.Done()
	for order := range orders {
		fmt.Printf("заказ ID: %d взял Worker %d\n", order.ID, id)
		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Симуляция времени обработки
		time.Sleep(time.Duration(rand.Intn(840)+30) * time.Millisecond) // случайное число 
		counters[id-1]++                                                // счетчик заказов для текущего воркера
	}
}

// генерирует заказы и закрывает канал
func GenerateOrders(orders chan<- Order, count int) {
	for i := 1; i <= count; i++ {
		order := Order{ID: i}
		fmt.Printf("заказ ID: %d создан\n", order.ID)
		orders <- order // Отправка заказа в буферизованный канал

		time.Sleep(time.Duration(rand.Intn(10)+10) * time.Millisecond) //  случайное число
	}
	close(orders) // Закрываем канал после генерации всех заказов
	fmt.Println("заказов больше нет.")
}

func main() {
	const numWorkers = 5
	const numOrders = 200

	orders := make(chan Order, 20)      // буферизованный канал заказов
	counters := make([]int, numWorkers) // массив счетчиков обработанных заказов по индексу worker-а

	var wg sync.WaitGroup

	// Запуск worker-ов
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, orders, &wg, counters)
	}

	time.Sleep(10 * time.Millisecond) // Задержка перед генерацией заказов....worker-ы ждут

	go GenerateOrders(orders, numOrders)

	wg.Wait() // Ждем завершения всех worker-ов
	fmt.Println("все завершились.")

	// статистика  количество обработанных заказов каждым вworker-ом
	var sum int = 0
	fmt.Printf("Worker -  orders.\n")
	for i := 0; i < numWorkers; i++ {
		fmt.Printf("# %d - %d.\n", i+1, counters[i])
		sum += counters[i]
	}
	fmt.Printf("всего заказов %d\n", sum)
}
