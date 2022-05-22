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
	fileName := "person.dat"
	writeData(fileName)
	readData(fileName)
}

func writeData(fileName string) {
	// начальные данные
	tom := person{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// сохраняем данные в файл
	fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
}

func readData(fileName string) {
	// переменная для считывания данных
	tom := person{}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// считывание данных из файла
	_, err = fmt.Fscanf(file, "%s %d %f\n", &tom.name, &tom.age, &tom.weight)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%-8s %-8d %-8.2f\n", tom.name, tom.age, tom.weight)
}
