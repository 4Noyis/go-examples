package main

import "fmt"

type rect struct {
	width, length int
}

/* Methods can be defined for either pointer or value receiver types. */
func (r *rect) area() int {
	return r.width * r.length
}

func (r rect) perim() int {
	return 2*r.width + 2*r.length
}

func main() {
	r := rect{10, 5}
	fmt.Println(r)

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	/* Go automatically handles conversion between values and pointers for method calls. */
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

}
