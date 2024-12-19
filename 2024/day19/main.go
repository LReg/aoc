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

func recSearch(combination string, patterns []string, currentPatterns []string) bool {
	joined := strings.Join(currentPatterns, "")
	if joined == combination {
		return true
	}
	if len(joined) >= len(combination) {
		return false
	}
	startC := combination[len(joined)]
	filteredPatterns := h.Filter(patterns, func(p string) bool {
		if p[0] == startC && len(combination) >= len(joined)+len(p) && p == combination[len(joined):len(joined)+len(p)] {
			return true
		} else {
			return false
		}
	})
	anyTrue := false
	for _, p := range filteredPatterns {
		if recSearch(combination, patterns, append(slices.Clone(currentPatterns), p)) {
			anyTrue = true
			break
		}
	}
	return anyTrue
}

func part1() {
	sum := 0
	lines := h.GetLinesAsSlice()
	patterns := []string{}
	combinations := []string{}
	stsec := 0
	for i, l := range lines {
		if len(l) == 0 {
			stsec = i + 1
			break
		}
		patt := strings.Split(l, ", ")
		patterns = append(patterns, patt...)
	}
	for _, l := range lines[stsec:] {
		combinations = append(combinations, l)
	}
	for i, c := range combinations {
		fmt.Println(i, "/", len(combinations))
		if recSearch(c, patterns, []string{}) {
			sum++
		}
	}

	fmt.Println(sum)
}

var cache = map[string]int{}

func recSearchp2(combination string, patterns []string, currentPatterns []string) int {
	joined := strings.Join(currentPatterns, "")
	r, ok := cache[joined]
	if ok {
		return r
	}
	if joined == combination {
		return 1
	}
	if len(joined) >= len(combination) {
		return 0
	}
	startC := combination[len(joined)]
	filteredPatterns := h.Filter(patterns, func(p string) bool {
		if p[0] == startC && len(combination) >= len(joined)+len(p) && p == combination[len(joined):len(joined)+len(p)] {
			return true
		} else {
			return false
		}
	})
	sumTrue := 0
	for _, p := range filteredPatterns {
		sumTrue += recSearchp2(combination, patterns, append(slices.Clone(currentPatterns), p))
	}
	cache[joined] = sumTrue
	return sumTrue
}

func part2() {
	sum := 0
	lines := h.GetLinesAsSlice()
	patterns := []string{}
	combinations := []string{}
	stsec := 0
	for i, l := range lines {
		if len(l) == 0 {
			stsec = i + 1
			break
		}
		patt := strings.Split(l, ", ")
		patterns = append(patterns, patt...)
	}
	for _, l := range lines[stsec:] {
		combinations = append(combinations, l)
	}
	for i, c := range combinations {
		fmt.Println(i, "/", len(combinations))
		sum += recSearchp2(c, patterns, []string{})
		cache = map[string]int{}
	}

	fmt.Println(sum)
}
