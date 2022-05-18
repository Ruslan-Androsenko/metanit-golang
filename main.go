package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")

	if err != nil { // если возникла ошибка
		fmt.Println(err)
		os.Exit(1) // выходим из программы
	}

	defer file.Close()
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)

		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}

		fmt.Print(string(data[:n]))
	}

	fmt.Println("\nDone.")
}
