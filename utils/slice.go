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

func GenerateCombinations[T any](input []T, length int) [][]T {
	// GenerateCombinations generates all combinations of the input slice with a given length.
	var result [][]T

	var helper func([]T)
	helper = func(current []T) {
		if len(current) == length {
			// Make a copy of the current slice and append to results
			temp := make([]T, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		// Recurse for each element in the input slice
		for _, val := range input {
			helper(append(current, val))
		}
	}

	helper([]T{})
	return result
}

func Map[T, V any](ts []T, fn func(T) (V, error)) []V {
	// applies fn to all elements of ts
	// works only for fn that returns an error as second element
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i], _ = fn(t)
	}
	return result
}
