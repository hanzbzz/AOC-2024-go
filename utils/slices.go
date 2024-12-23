package utils

func CountOccurrences[S ~[]E, E comparable](input S) map[E]int {
	// Count the number of occurrences of a value in a slice.
	result := map[E]int{}
	for _, v := range input {
		result[v]++
	}
	return result
}
