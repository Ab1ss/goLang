package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	sum := 0
	for _, num := range arr1 {
		sum += num * num
	}
	fmt.Println(sum)
}
