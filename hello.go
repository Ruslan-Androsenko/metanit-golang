package main

import "fmt"

type library []string

func main() {
	var lib library = library{"Book1", "Book2", "Book3"}
	lib.print()
}

func (l library) print() {
	for _, val := range l {
		fmt.Println(val)
	}
}
