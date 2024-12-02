package main

import (
	"AOC/helper"
	"fmt"
)

func main() {
	lines := helper.GetLinesAsSlice()
	sum := 0
	for _, l := range lines {
		nums := helper.ExtrapolateNumbersFromString(l, " ")
		valid := true
		inc := nums[0] < nums[1]
		for i := 0; i+1 < len(nums); i++ {
			rawDiff := nums[i] - nums[i+1]
			diff := helper.Abs(nums[i] - nums[i+1])

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
