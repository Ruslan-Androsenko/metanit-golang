package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte("Hello Bold!")
	file, err := os.Create("hello.bin")

	if err != nil { // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}

	defer file.Close()
	file.Write(data)

	fmt.Println("Done.")
}
