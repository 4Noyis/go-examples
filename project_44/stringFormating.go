package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	/*If the value is a struct, the %+v variant will include the struct’s field names.*/
	fmt.Printf("struct2: %+v\n", p)

	/*The %#v variant prints a Go syntax representation of the value*/
	fmt.Printf("struct3: %#v\n", p)

	/*To print the type of a value, use %T.*/
	fmt.Printf("type: %T\n", p)

	fmt.Printf("bool: %t\n", true)

	/*Use %d for standard, base-10 formatting.*/
	fmt.Printf("int: %d\n", 123)

	/*This prints a binary representation.*/
	fmt.Printf("bin: %b\n", 14)

	/*This prints the character corresponding to the given integer.*/
	fmt.Printf("char: %c\n", 33)

	/*%x provides hex encoding.*/
	fmt.Printf("hex: %x\n", 456)

	/*For basic decimal formatting use %f.*/
	fmt.Printf("float1: %f\n", 78.9)

	/*%e and %E format the float in (slightly different versions of) scientific notation.*/
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	/*  For basic string printing use %s*/
	fmt.Printf("str1: %s\n", "\"string\"")

	/*As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.*/
	fmt.Printf("str3: %x\n", "hex this")

	/*To print a representation of a pointer, use %p.*/
	fmt.Printf("pointer: %p\n", &p)

	/*When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb.*/
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	/* So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.*/
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	/*You can format+print to io.Writers other than os.Stdout using Fprintf.*/
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
