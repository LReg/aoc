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
	for _, l := range lines {
		ns := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)
		if ns[0] <= ns[2] && ns[1] >= ns[3] {
			sum++
		} else if ns[0] >= ns[2] && ns[1] <= ns[3] {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		ns := h.ExtrapolateNumbersFromStringIgnoreNonDig(l)
		if ns[1] >= ns[2] && ns[0] <= ns[3] {
			sum++
		}
	}
	fmt.Println(sum)
}
