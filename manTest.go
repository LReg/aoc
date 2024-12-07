package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	elements := []string{"ADD", "MULT"}
	n := 2
	result := h.CrossProduct(elements, n)

	// Ergebnis ausgeben
	for _, combination := range result {
		fmt.Println(combination)
	}
}
