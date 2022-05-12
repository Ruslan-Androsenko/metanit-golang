package main

import "fmt"

func main() {
	// Empty pointer
	var pf *float64
	fmt.Println("Address:", pf)

	if pf != nil {
		fmt.Println("Value:", *pf)
	}

	// Dynamic object
	p := new(int)
	fmt.Println("value:", *p)

	*p = 8
	fmt.Println("value:", *p)
}
