package main

import (
	"AOC/h"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

func part1() {
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		b := []byte(l)
		for i, _ := range b {
			sl := make([]byte, 0)
			sl = append(sl, b[i])
			good := true
			for _, c := range h.R(1, 4) {
				if slices.Contains(sl, b[i+c]) {
					good = false
					break
				} else {
					sl = append(sl, b[i+c])
				}
			}
			if good {
				fmt.Println(i + 4)
				break
			}
		}
	}
}

func part2() {
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		b := []byte(l)
		for i, _ := range b {
			sl := make([]byte, 0)
			sl = append(sl, b[i])
			good := true
			for _, c := range h.R(1, 14) {
				if slices.Contains(sl, b[i+c]) {
					good = false
					break
				} else {
					sl = append(sl, b[i+c])
				}
			}
			if good {
				fmt.Println(i + 14)
				break
			}
		}
	}
}
