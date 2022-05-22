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
	tom := person{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}

	file, err := os.Create("person.dat")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Fprintf(file, "%-10s %-10d %-10.3f\n", tom.name, tom.age, tom.weight)
}
