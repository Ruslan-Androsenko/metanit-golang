package main

import (
	"fmt"
	"hello/computation"
	"rsc.io/quote"
)

func main() {
	message := quote.Hello()
	fmt.Println(message)

	fmt.Println()
	fmt.Println(computation.Factorial(4))
	fmt.Println(computation.Factorial(5))
}
