package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 10
	fmt.Println("length:", len(intCh)) // 1
	intCh <- 3
	intCh <- 24
	//intCh <- 15	// блокировка - функций main ждет, когда освободится место в канале

	fmt.Println("\nintCh:", <-intCh) // 10
	fmt.Println("intCh:", <-intCh)   // 3
	fmt.Println("intCh:", <-intCh)   // 24

	fmt.Println("\ncapacity:", cap(intCh)) // 3
	fmt.Println("length:", len(intCh))
}

func factorial(n int, ch chan int) {
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println(n, "-", result)
	ch <- result // отправка данных в канал
}
