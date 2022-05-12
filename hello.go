package main

import "fmt"

func main() {
	var x int = 4
	var p *int

	p = &x
	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)

	*p = 25
	fmt.Println(x)

	f := 2.3
	pf := &f

	fmt.Println("Address:", pf)
	fmt.Println("Value:", *pf)
}
