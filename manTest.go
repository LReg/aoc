package main

import "fmt"

func main() {
	a := []int{1, 1, 1, 1}
	b := append(a, 2)
	c := append(a, 3)
	a[0] = 0
	fmt.Println(a, b, c)
}
