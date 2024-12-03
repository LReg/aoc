package main

import "strings"

func main() {
	s := "123456789123456789"
	a := strings.Index(s, "7")
	b := strings.Index(s[10:], "7")
	println(a, b)
}
