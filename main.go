package main

import "fmt"

func main() {
	//var inCh chan<- int  // Канал только для отправки данных
	//var outCh <-chan int // Канал только для получения данных

	intCh := make(chan int, 2)
	go factorial(5, intCh)

	fmt.Println("intCh:", <-intCh)
	fmt.Println("The End")
}

func factorial(n int, ch chan<- int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
	ch <- result // отправка данных в канал
}
