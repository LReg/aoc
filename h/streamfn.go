package h

import (
	"strconv"
	"strings"
)

func MapStrToInt(s string) int {
	s = strings.TrimSpace(s)
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func ReduceSumInt(prev int, curr int) int {
	return prev + curr
}

func MapReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
