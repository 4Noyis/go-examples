package main

import (
	"fmt"
)

/*
Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).
*/
func main() {
	m := make(map[string]int) /*make(map[key-type]val-type)*/

	m["k1"] = 3
	m["k2"] = 5

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("value 1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k1")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
