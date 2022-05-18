package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	close(intCh) // канал закрыт

	//intCh <- 24 // ошибка - канал уже закрыт

	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened {
			fmt.Println("val:", val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}
