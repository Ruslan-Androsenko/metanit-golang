package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("борщ с капустой, но не красный")
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}
