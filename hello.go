package main

import "fmt"

func main() {
	f := selectFn(1)
	fmt.Println(f(3, 4)) // 7

	f = selectFn(3)
	fmt.Println(f(3, 4)) // 12

	f = selectFn(2)
	fmt.Println(f(4, 3)) // 1
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

func isEven(n int) bool {
	return n%2 == 0
}

func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {
	result := 0

	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}

	return result
}
