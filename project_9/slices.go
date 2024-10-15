package main

import (
	"fmt"
	"slices"
)

func main() {

	s := make([]string, 3)
	fmt.Println("emo:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "e", "f")
	fmt.Println("appended: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copyed:", c)

	t := []int8{1, 2, 3}
	t2 := []int8{1, 2, 3}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 is equal", t)
	}
}
