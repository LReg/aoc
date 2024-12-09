package h

import "cmp"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min[T cmp.Ordered](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[T cmp.Ordered](a T, b T) T {
	if a < b {
		return b
	} else {
		return a
	}
}
