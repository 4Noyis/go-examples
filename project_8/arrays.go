package main

import "fmt"

func main() {
	/*
		  In Go, an array is a numbered sequence of elements of a specific length.
		In typical Go code, slices are much more common; arrays are useful in some special scenarios
	*/
	var array [6]int
	fmt.Println("emp:", array)

	array[4] = 100
	fmt.Println("set", array)
	fmt.Println("get", array[4])

	fmt.Println("length", len(array))

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
