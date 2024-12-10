package main

import (
	"AOC/h"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

const (
	ADD = iota
	MUL
	CONCAT
)

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		parts := strings.Split(l, ":")
		res, _ := strconv.Atoi(parts[0])
		vals := h.ExtrapolateNumbersFromString(parts[1], " ")
		cp := h.Permutations([]int{ADD, MUL, CONCAT}, len(vals)-1)
		for _, p := range cp {
			acRes := vals[0]
			for i, operation := range p {
				if operation == CONCAT {
					r, _ := strconv.Atoi(strconv.Itoa(acRes) + strconv.Itoa(vals[i+1]))
					acRes = r
				} else if operation == ADD {
					acRes = acRes + vals[i+1]
				} else {
					acRes = acRes * vals[i+1]
				}
			}
			if acRes == res {
				sum += res
				break
			}
		}
	}
	fmt.Println(sum)
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		parts := strings.Split(l, ":")
		res, _ := strconv.Atoi(parts[0])
		vals := h.ExtrapolateNumbersFromString(parts[1], " ")
		cp := h.Permutations([]int{ADD, MUL}, len(vals)-1)
		for _, p := range cp {
			acRes := vals[0]
			for i, operation := range p {
				if operation == ADD {
					acRes = acRes + vals[i+1]
				} else {
					acRes = acRes * vals[i+1]
				}
			}
			if acRes == res {
				sum += res
				break
			}
		}
	}
	fmt.Println(sum)
}
