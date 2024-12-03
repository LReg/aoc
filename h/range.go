package h

// R Range from f to t; t excluding
func R(f int, t int) []int {
	r := make([]int, 0)
	for ; f < t; f++ {
		r = append(r, f)
	}
	return r
}

// Iter Range Sequence; RS(2) -> [0, 1]
func Iter(t int) []int {
	r := make([]int, 0)
	for i := 0; i < t; i++ {
		r = append(r, i)
	}
	return r
}
