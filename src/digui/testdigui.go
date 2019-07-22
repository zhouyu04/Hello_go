package main

import (
	"fmt"
)

func Factorial(n uint16) (result uint16) {

	if n > 0 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	result := Factorial(5)

	fmt.Println("result:", result)
}
