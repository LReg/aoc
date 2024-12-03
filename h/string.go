package h

import (
	"strconv"
	"strings"
	"unicode"
)

func ExtrapolateNumbersFromString(s string, seperator string) []int {
	nums := make([]int, 0, 50)
	numstrs := strings.Split(s, seperator)
	for _, s := range numstrs {
		s = strings.TrimSpace(s)
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

func ExtrapolateNumbersFromStringIgnore(s string, seperator string, ignore []string) []int {
	for _, i := range ignore {
		s = strings.ReplaceAll(s, i, "")
	}
	return ExtrapolateNumbersFromString(s, seperator)
}

func ExtrapolateNumbersFromStringIgnoreNonDig(s string) []int {
	return ExtrapolateNumbersFromString(RemoveNonDigits(s), " ")
}

func RemoveNonDigits(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if !unicode.IsDigit(runes[i]) {
			runes[i] = ' '
		}
	}
	return string(runes)
}

func DelChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
