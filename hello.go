package main

import "fmt"

type mile int
type kilometer int
type library []string

func main() {
	fmt.Println("Type mile")
	var distance mile = 5
	distanceToEnemy(distance)

	//var distance2 kilometer = 5
	//distanceToEnemy(distance2)	! ошибка, передан аргумент другого типа

	fmt.Println("\nType library")
	var myLibrary library = library{"Book1", "Book2", "Book3"}
	printBooks(myLibrary)
}

func distanceToEnemy(distance mile) {
	fmt.Println("Расстояние до противника:")
	fmt.Println(distance, "миль")
}

func printBooks(lib library) {
	for _, value := range lib {
		fmt.Println(value)
	}
}
