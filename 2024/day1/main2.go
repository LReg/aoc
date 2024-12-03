package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	lines := h.GetLinesAsSlice()
	left := make([]int, 0)
	right := make([]int, 0)
	for _, l := range lines {
		nums := h.ExtrapolateNumbersFromString(l, " ")
		left = append(left, nums[0])
		right = append(right, nums[1])
	}

	m := make(map[int]int)
	for _, r := range right {
		m[r] += 1
	}
	sum := 0
	for _, l := range left {
		n, ok := m[l]
		if !ok {
			continue
		}
		sum += l * n
	}

	fmt.Print(sum)
}
