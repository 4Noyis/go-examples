package main

/*
Sometimes we’ll want to sort a collection by something other than its natural order.
For example, suppose we wanted to sort strings by their length instead of alphabetically*/
import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lencmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	/*  Call slices.SortFunc with this custom comparison function to sort fruits by name length.*/
	slices.SortFunc(fruits, lencmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)

}
