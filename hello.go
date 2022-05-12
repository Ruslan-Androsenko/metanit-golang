package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue
	fmt.Println(colors[1]) // green
}
