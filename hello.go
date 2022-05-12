package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var tom = person{name: "Tom", age: 24}
	var tomPointer *person = &tom
	fmt.Println("before:", tom.age)

	tomPointer.updateAge(33)
	fmt.Println("after:", tom.age)
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}
