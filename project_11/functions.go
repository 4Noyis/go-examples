package main

import "fmt"

func plus(a int, b int) int {
	sum := a + b
	return sum
}

func plusPlus(a, b, c int) int {

	return a + b + c
}

func main() {
	result := plus(10, 25)
	fmt.Println(result)

	result = plusPlus(3, 5, 2)
	fmt.Println(result)
}
