package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, c := range lines[0] {
		if c == '(' {
			sum++
		} else {
			sum--
		}
	}
	fmt.Println(sum)
}

func part2() {

	sum := 0
	lines := h.GetLinesAsSlice()
	for i, c := range lines[0] {
		if c == '(' {
			sum++
		} else {
			sum--
		}
		if sum == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
