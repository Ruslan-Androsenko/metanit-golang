package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	fmt.Println(people)
	fmt.Println(people["Alice"]) // 8
	fmt.Println(people["Bob"])   // 2

	people["Bob"] = 32
	fmt.Println(people["Bob"]) // 32

	fmt.Println("\nCheck item by key")
	// Проверка наличия элемента по определенному ключу
	if val, ok := people["Tom"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Item is not found.")
	}

	fmt.Println("\nForeach for map.")
	for key, value := range people {
		fmt.Println(key, value)
	}
}
