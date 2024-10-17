package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Sorted String: ", strs)

	ints := []int16{5, 1, 4, 9}
	slices.Sort(ints)
	fmt.Println("Sorted Int: ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
