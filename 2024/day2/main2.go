package main

import (
	"AOC/helper"
	"fmt"
)

func main() {
	fmt.Println("Start")
	lines := helper.GetLinesAsSlice()
	sum := 0
	for _, l := range lines {
		nums := helper.ExtrapolateNumbersFromString(l, " ")
		valid, _ := isValid(nums)
		if valid {
			sum++
			continue
		} else {
			v := false
			for i := 0; i < len(nums); i++ {
				copyNums := make([]int, len(nums)-1)
				copy(copyNums, nums[:i])
				copy(copyNums[i:], nums[i+1:])
				va, _ := isValid(copyNums)
				if va {
					v = true
					break
				}
			}
			if v {
				sum++
			}
		}
	}
	fmt.Printf("Sum: %d", sum)
}

func isValid(nums []int) (bool, int) {
	inc := nums[0] < nums[1]
	for i := 0; i+1 < len(nums); i++ {
		rawDiff := nums[i] - nums[i+1]
		diff := helper.Abs(nums[i] - nums[i+1])

		if diff < 1 || diff > 3 {
			return false, i
		} else if rawDiff > 0 && inc {
			return false, i
		} else if rawDiff < 0 && !inc {
			return false, i
		}
	}
	return true, -1
}
