package main

import (
	"fmt"

	kata "github.com/Ab1ss/goLang/codeWars"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := kata.SquareSum(numbers)
	fmt.Println(result)
}
