package main

import "fmt"

func main() {
	fmt.Println("sum")
	f := selectFn(1)
	fmt.Println(f(2, 3)) // 5
	fmt.Println(f(4, 5)) // 9
	fmt.Println(f(6, 7)) // 13

	fmt.Println("\nsubtraction")
	f = selectFn(2)
	fmt.Println(f(2, 3)) // -1
	fmt.Println(f(7, 4)) // 3

	fmt.Println("\nmultiply")
	f = selectFn(5)
	fmt.Println(f(2, 3)) // 6
	fmt.Println(f(4, 5)) // 20
	fmt.Println(f(6, 7)) // 42
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}
