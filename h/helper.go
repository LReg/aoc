package h

func DeleteIndexFromSlice[T any](sl []T, index int) []T {
	return append(sl[:index], sl[index+1:]...)
}
