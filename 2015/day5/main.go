package main

import (
	"AOC/h"
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Start")
	part1()
	part2()
}

var v = []byte("aeiou")
var no = []string{"ab", "cd", "pq", "xy"}

func has3Vowls(s string) bool {
	c := 0
	for _, b := range s {
		if slices.Contains(v, byte(b)) {
			c++
		}
	}
	if c >= 3 {
		return true
	}
	return false
}

func hasTwiceInARow(s string) bool {
	prev := 'Y'
	for _, r := range s {
		if prev == r {
			return true
		} else {
			prev = r
		}
	}
	return false
}

func doesNotContainNo(s string) bool {
	for _, n := range no {
		if strings.Contains(s, n) {
			return false
		}
	}
	return true
}

func hasRepeatingPair(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		p := h.Pair[byte]{s[i], s[i+1]}
		for j := i + 2; j < len(s)-1; j++ {
			sp := h.Pair[byte]{s[j], s[j+1]}
			if p == sp {
				return true
			}
		}
	}
	return false
}

func rep1letApart(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		if has3Vowls(l) && hasTwiceInARow(l) && doesNotContainNo(l) {
			sum++
		}
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	for _, l := range lines {
		if hasRepeatingPair(l) && rep1letApart(l) {
			sum++
		}
	}
	fmt.Println(sum)
}
