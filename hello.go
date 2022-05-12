package main

import "fmt"

func main() {
	var a = add(4, 5)
	var b = add(20, 6)

	fmt.Println(a)
	fmt.Println(b)

	var age, name = add3(4, 5, "Tom", "Simpson")
	fmt.Println("\nage: ", age, "\nname: ", name)
}

func add(x, y int) int {
	return x + y
}

func add2(x, y int) (z int) {
	z = x + y
	return
}

func add3(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName

	return z, fullName
}
