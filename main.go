package main

import (
	"fmt"
	"io"
	"os"
)

type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0

	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}

	return count, io.EOF
}

func main() {
	file, err := os.Open("hello.txt")

	if err != nil { // если возникла ошибка
		fmt.Println(err)
		os.Exit(1) // выходим из программы
	}

	defer file.Close()

	io.Copy(os.Stdout, file)
	fmt.Println()

	phone1 := phoneReader("+1(234)567 90-10")
	io.Copy(os.Stdout, phone1)

	fmt.Println("\nDone.")
}
