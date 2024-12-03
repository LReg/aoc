package h

func Map[T any, U any](input []T, f func(T) U) []U {
	// Erstellt ein Ergebnis-Slice der gewünschten Typen U
	result := make([]U, len(input))
	// Wendet die Funktion auf jedes Element des Input-Slices an
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any](input []T, f func(T, T) T, initial T) T {
	result := initial
	for _, v := range input {
		result = f(result, v)
	}
	return result
}
