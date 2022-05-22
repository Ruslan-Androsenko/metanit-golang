package main

import (
	"fmt"
	"io"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	fileName := "people.dat"
	writeData(fileName)
	readData(fileName)
}

func writeData(fileName string) {
	// начальные данные
	var people = []person{
		{name: "Tom", age: 24, weight: 68.5},
		{name: "Bob", age: 25, weight: 64.2},
		{name: "Sam", age: 27, weight: 73.6},
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// сохраняем данные в файл
	for _, p := range people {
		fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
	}
}

func readData(fileName string) {
	// переменные для считывания данных
	var name string
	var age int
	var weight float64

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	// считывание данных из файла
	for {
		_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}
}
