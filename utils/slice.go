package utils

func CountOccurrences[S ~[]E, E comparable](input S) map[E]int {
	// Count the number of occurrences of a value in a slice.
	result := map[E]int{}
	for _, v := range input {
		result[v]++
	}
	return result
}

func RemoveIndex[S any](input []S, index int) []S {
	// Remove an element from a slice by index.
	return append(input[:index], input[index+1:]...)
}
