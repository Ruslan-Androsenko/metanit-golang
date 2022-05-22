package main

import (
	"fmt"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	fileName := "hello2.txt"
	writeData(fileName)
	readData(fileName)
}

func writeData(fileName string) {
	// начальные данные
	var name string = "Tom"
	var age int = 24

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}

func readData(fileName string) {
	var name string
	var age int

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}
