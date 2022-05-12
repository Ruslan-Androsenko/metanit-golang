package main

import "fmt"

func main() {
	hello()

	add(4, 5, 3.4, 5.6, 1.2)
	add(20, 6, 3.14, 9.81, 5.76)

	var a = 8
	fmt.Println("\na before: ", a)
	increment(a)
	fmt.Println("a before: ", a)

	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	sum(5, 6, 7, 2, 3)

	var nums = []int{5, 6, 7, 2, 3}
	sum(nums...)
}

func hello() {
	fmt.Println("Hello World!")
}

func add(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c

	fmt.Println("\nx + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func sum(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println("\nsum: ", sum)
}

func increment(x int) {
	fmt.Println("x before: ", x)
	x += 20
	fmt.Println("x after: ", x)
}
