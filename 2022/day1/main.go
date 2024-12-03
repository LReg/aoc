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
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1])
}
