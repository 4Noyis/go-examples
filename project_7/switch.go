package main

import (
	"fmt"

	"time"
)

func main() {

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend")
	default:
		fmt.Println("Its Weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its Before noon")
	default:
		fmt.Println("Its After noon")
	}

	/*
	  A type switch compares types instead of values.
	  You can use this to discover the type of an interface value.
	*/
	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am Int")
		default:
			fmt.Printf("I dont know type %T\n", t)
		}
	}
	WhatAmI(true)
	WhatAmI(12)
	WhatAmI("hey")
}
