package main

import "fmt"

func main() {
	d := 5
	fmt.Println("d before:", d)

	changeValue(&d)
	fmt.Println("d after:", d)
}

func changeValue(x *int) {
	*x = (*x) * (*x)
}
