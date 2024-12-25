package utils

func Apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}
