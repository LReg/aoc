package helper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetLinesAsSlice() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	slice := make([]string, 0, 50)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	return slice
}

func GetLinesAsRuneSlices() [][]rune {
	lines := GetLinesAsSlice()
	lr := make([][]rune, 0)
	for _, l := range lines {
		ra := []rune(l)
		lr = append(lr, ra)
	}
	return lr
}

func ExtrapolateNumbersFromString(s string, seperator string) []int {
	nums := make([]int, 0, 50)
	numstrs := strings.Split(s, seperator)
	for _, s := range numstrs {
		s = strings.ReplaceAll(s, " ", "")
		if len(s) > 0 {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
	}
	return nums
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
