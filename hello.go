package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")

	var (
		name string = "Tom"
		age  int    = 27
	)

	fmt.Println("name:", name) // Tom
	fmt.Println("age:", age)   // 27

	const pi float64 = 3.1415
	fmt.Println("pi:", pi)

	age++
	fmt.Println("age:", age)

	var nameTwo string = "Pol"
	var ageTwo int = 32

	fmt.Println("nameTwo:", nameTwo)
	fmt.Println("ageTwo:", ageTwo)
}
