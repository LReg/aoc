package h

// should probably not be used since I found the slices.xxx() functions

func Push[T any](stack *[]T, e T) {
	*stack = append(*stack, e)
}

func Pop[T any](stack *[]T) T {
	var r T
	if len(*stack) == 0 {
		return r
	}
	r = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return r
}

func Peek[T any](stack *[]T) T {
	var r T
	if len(*stack) == 0 {
		return r
	}
	return (*stack)[len(*stack)-1]
}
