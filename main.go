package main

import "fmt"

func main() {
	var intCh chan int = make(chan int) // канал для данных типа int
	strCh := make(chan string)          // канал для данных типа string

	intCh <- 5     // отправка данных  в канал
	val := <-intCh // получение данных из канала

	fmt.Println("val:", val)
	fmt.Println("intCh:", intCh)
	fmt.Println("strCh:", strCh)
}
