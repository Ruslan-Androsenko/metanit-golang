package main

import "fmt"

func main() {
	p1 := createPointer(7)
	fmt.Println("p1:", *p1)

	p2 := createPointer(10)
	fmt.Println("p1:", *p2)

	p3 := createPointer(28)
	fmt.Println("p1:", *p3)
}

func createPointer(x int) *int {
	p := new(int)
	*p = x

	return p
}
