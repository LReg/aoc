package main

import (
	"AOC/h"
	"fmt"
	"strings"
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
		a := strings.Split(l, ":")
		password := strings.TrimSpace(a[1])
		b := strings.Split(a[0], " ")
		fromto := h.ExtrapolateNumbersFromString(b[0], "-")
		contain := b[1]
		contains := 0
		for _, c := range password {
			if string(c) == contain {
				contains++
			}
		}
		//fmt.Println(password, fromto, contain, contains)
		if fromto[0] <= contains && fromto[1] >= contains {
			sum++
		}
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		a := strings.Split(l, ":")
		password := strings.TrimSpace(a[1])
		b := strings.Split(a[0], " ")
		fromto := h.ExtrapolateNumbersFromString(b[0], "-")
		contain := b[1]
		//fmt.Println(password, fromto, contain, contains)
		if string(password[fromto[0]-1]) == contain && !(string(password[fromto[1]-1]) == contain) {
			sum++
			continue
		} else if !(string(password[fromto[0]-1]) == contain) && string(password[fromto[1]-1]) == contain {
			sum++
			continue
		}
	}

	fmt.Println(sum)
}
