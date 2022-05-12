package main

import "fmt"

func main() {
	// var f func(int, int) int = add
	f := add
	fmt.Println("sum:", f(3, 4))

	f = multiply
	fmt.Println("multiply:", f(3, 4))

	var f1 func(string) = display
	f1("hello")
}

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func display(message string) {
	fmt.Println(message)
}
