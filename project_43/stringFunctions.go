package main

import (
	"fmt"
	s "strings"
)

/*We alias fmt.Println to a shorter name as we’ll use it a lot below.*/
var p = fmt.Println

func main() {

	p("contains:", s.Contains("test", "es"))
	p("count: ", s.Contains("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat: ", s.Repeat("a", 6))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

}
