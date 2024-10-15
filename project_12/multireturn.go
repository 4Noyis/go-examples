package main

import "fmt"

func vals() (int, int) {

	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println("val1: ", a, "val2: ", b)

	_, a = vals()
	fmt.Println("a: ", a, "b: ", b)
}
