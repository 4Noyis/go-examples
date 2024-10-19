package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	/*Here we start with a new Hash*/
	h := sha256.New()

	/*Write expects bytes. If you have a string s, use []byte(s) to coerce it to bytes.*/
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
