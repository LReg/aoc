package main

import (
	"AOC/h"
	"fmt"
)

func main() {
	lines := h.GetLinesAsSlice()
	sum := 0
	for _, l := range lines {
		nums := h.ExtrapolateNumbersFromString(l, " ")
		valid := true
		inc := nums[0] < nums[1]
		for i := 0; i+1 < len(nums); i++ {
			rawDiff := nums[i] - nums[i+1]
			diff := h.Abs(nums[i] - nums[i+1])

			if diff < 1 || diff > 3 {
				valid = false
				break
			}
			if rawDiff > 0 && inc {
				valid = false
				break
			} else if rawDiff < 0 && !inc {
				valid = false
				break
			}
		}
		if valid == true {
			sum++
		}
	}
	fmt.Printf("Sum: %d", sum)
}
