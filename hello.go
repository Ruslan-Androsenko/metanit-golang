package main

import "fmt"

func main() {
	f := square()
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // 25
}

func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}
