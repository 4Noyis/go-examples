package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)

	/*User contains all authentication info; call Username and Password on this for individual values.*/
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println("-----------------------------------")
	/*The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.*/
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println("-----------------------------------")

	/*Here we extract the path and the fragment after the #.*/
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println("-----------------------------------")

	/*To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map.
	The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.*/
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
