package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//var tom person = person {"Tom", 24}
	var alice person = person{age: 23, name: "Alice"}
	var tom = person{name: "Tom", age: 24}
	bob := person{name: "Bob", age: 31}

	fmt.Println(alice)
	fmt.Println(tom)
	fmt.Println(bob)
}
