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
	for i, l := range lines {
		num := h.ExtrapolateNumbersFromString(l, " ")[0]
		for _, ll := range lines[i:] {
			numm := h.ExtrapolateNumbersFromString(ll, " ")[0]
			if num+numm == 2020 {
				fmt.Println(num * numm)
				return
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for i, l := range lines {
		num := h.ExtrapolateNumbersFromString(l, " ")[0]
		for j, ll := range lines[i:] {
			numm := h.ExtrapolateNumbersFromString(ll, " ")[0]
			for _, lll := range lines[j:] {
				nummm := h.ExtrapolateNumbersFromString(lll, " ")[0]
				if num+numm+nummm == 2020 {
					fmt.Println(num * numm * nummm)
					return
				}
			}
		}
	}

	fmt.Println(sum)
}
