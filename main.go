package main

import "fmt"

func main() {
	fmt.Println("Start")

	// Создание канала и получение из него данных
	fmt.Println(<-createChan(5)) // 5
	fmt.Println("The End")
}

func createChan(n int) chan int {
	ch := make(chan int) // создаем канал

	go func() {
		ch <- n // отправляем данные в канал
	}() // запускаем горутину

	return ch // возврщаем канал
}

func factorial(n int, ch chan<- int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
	ch <- result // отправка данных в канал
}
