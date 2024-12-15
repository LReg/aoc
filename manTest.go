package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	pq := h.NewPC[string]()
	pq.Push("23", 2)
	pq.Push("345", 5)
	pq.Push("-101", 1)
	fmt.Println(pq.First())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
}
