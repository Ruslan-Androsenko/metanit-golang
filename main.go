package main

import "fmt"

func main() {
	intCh := make(chan int)

	/*
		go func() {
			fmt.Println("Go routine starts")
			intCh <- 5 // блокировка, пока данные не будут получены функцией main
		}()
	*/

	go factorial(5, intCh) // вызов горутины

	fmt.Println("intCh:", <-intCh) // получение данных из канала
	fmt.Println("The End")

	//intCh2 := make(chan int)
	//intCh2 <- 10	// функция main блокируется
	//fmt.Println("intCh2:", <-intCh2)
}

func factorial(n int, ch chan int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
	ch <- result // отправка данных в канал
}
