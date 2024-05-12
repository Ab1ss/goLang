package main

import (
	"fmt"

	"codeWars/kata" // Import the kata package from test.go
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := kata.SquareSum(numbers)
	fmt.Println(result)
}
