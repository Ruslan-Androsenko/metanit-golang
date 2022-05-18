package main

import "fmt"

func main() {
	intCh := make(chan int)
	go factorial(7, intCh)

	/*
		// получение данных из потока с помощью бесконечного цикла
		for {
			// получаем данные из потока
			if num, opened := <-intCh; opened {
				fmt.Println(num)
			} else {
				break // если поток закрыт, выход из цикла
			}
		}
	*/

	// получение данных из потока с помощью другой формы цикла
	for num := range intCh {
		fmt.Println(num)
	}
}

func factorial(n int, ch chan int) {
	defer close(ch)
	result := 1

	for i := 1; i <= n; i++ {
		result *= i
		ch <- result // посылаем по числу
	}
}
