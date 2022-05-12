package main

import "fmt"

func main() {
	fmt.Println("factorial")
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720

	fmt.Println("\nfibbonachi")
	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 4
	fmt.Println(fibbonachi(6)) // 8
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func fibbonachi(n uint) uint {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibbonachi(n-1) + fibbonachi(n-2)
}
