package main

import (
	"AOC/helper"
	"fmt"
	"sort"
)

func main() {
	lines := helper.GetLinesAsSlice()
	left := make([]int, 0)
	right := make([]int, 0)
	for _, l := range lines {
		nums := helper.ExtrapolateNumbersFromString(l, " ")
		left = append(left, nums[0])
		right = append(right, nums[1])
	}
	sort.IntSlice(left).Sort()
	sort.IntSlice(right).Sort()
	sum := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Print(sum)
}
