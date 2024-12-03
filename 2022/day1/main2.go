package main

import (
	"AOC/h"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := h.GetLinesAsSlice()
	elves := make([]int, 0)
	cals := 0
	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			elves = append(elves, cals)
			cals = 0
			continue
		}
		cals += n
	}
	elves = append(elves, cals)
	sort.Ints(elves)
	t3 := make([]int, 3)
	copy(t3, elves[len(elves)-3:])
	sum := 0
	for _, c := range t3 {
		sum += c
	}
	fmt.Println(sum)
}
