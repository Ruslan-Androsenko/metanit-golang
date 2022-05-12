package main

import "fmt"

func main() {
	var users = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users)
	fmt.Println(users[2]) // Kate

	fmt.Println("\nChange item value")
	users[2] = "Katherine"

	for _, value := range users {
		fmt.Println(value)
	}

	fmt.Println("\nFunction make")
	var usersTest []string = make([]string, 3)
	usersTest[0] = "Tom"
	usersTest[1] = "Alice"
	usersTest[2] = "Bob"
	fmt.Println(usersTest)

	fmt.Println("\nAdding items to slice")
	users = append(users, "Bob")

	for _, value := range users {
		fmt.Println(value)
	}

	fmt.Println("\nOperator the slice")
	// Исходный массив
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 0-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1)
	fmt.Println(users2)
	fmt.Println(users3)

	fmt.Println("\nRemove item from slice")
	var n = 3
	users = append(initialUsers[:n], initialUsers[n+1:]...)
	fmt.Println(users)
}
