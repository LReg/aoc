package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{7, 4, 5}
	fmt.Println(slices.Min(s))
}
