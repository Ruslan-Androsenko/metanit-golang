package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{name: "Tom", age: 22}

	var tomPointer *person = &tom
	tomPointer.age = 29
	fmt.Println(tom.age)

	(*tomPointer).age = 32
	fmt.Println(tom.age)

	fmt.Println("\nУказатели на безымянные объекты")
	var alice *person = &person{name: "Alice", age: 23}
	var bob *person = new(person)

	fmt.Println(alice)
	fmt.Println(bob)

	fmt.Println("\nУказатель на поле объекта")
	var agePointer *int = &tom.age

	*agePointer = 35
	fmt.Println(tom.age)
}
