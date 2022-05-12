package main

import "fmt"

func main() {
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var (
		sum  = 0
		prod = 1
	)

	var i = 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	fmt.Println()
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	fmt.Println()
	for index, item := range numbers {
		fmt.Println(index, ").", item)
	}

	for _, item := range numbers {
		if item < 0 {
			continue
		}

		sum += item
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] != 0 {
			prod *= numbers[i]
		}
	}

	fmt.Println("\nsum: ", sum)
	fmt.Println("prod: ", prod)
}
