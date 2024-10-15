package main

import "fmt"

func main() {
	fmt.Println("z")

	if number := 100; number < 0 {
		fmt.Println(number, "is negative ")
	} else if number < 10 {
		fmt.Println(number, "has one digit")
	} else {
		fmt.Println(number, "has multiple digit")
	}
}
