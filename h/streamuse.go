package h

import (
	"strconv"
	"strings"
)

// StrSlToIntSlSoftFail slice, failed
func StrSlToIntSlSoftFail(sl []string) ([]int, bool) {
	failed := false
	nums := Map(sl, func(s string) int {
		s = strings.TrimSpace(s)
		atoi, err := strconv.Atoi(s)
		if err != nil {
			failed = true
		}
		return atoi
	})
	return nums, failed
}

func StrSlToIntSl(sl []string) []int {
	return Map(sl, MapStrToInt)
}

func SumIntSl(sl []int) int {
	return Reduce(sl, ReduceSumInt, 0)
}
