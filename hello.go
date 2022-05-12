package main

import "fmt"

func main() {
	action(10, 25, add)    // 35
	action(5, 6, multiply) // 30

	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEvens := sum(slice, isEven) // Сумма четных чисел
	fmt.Println(sumOfEvens)          // -2

	sumOfPositive := sum(slice, isPositive) // Сумма положительных чисел
	fmt.Println(sumOfPositive)              // 37
}

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
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
