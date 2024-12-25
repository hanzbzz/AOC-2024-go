package utils

import (
	"os"
	"strings"
)

func ReadFileLines(path string) ([]string, error) {
	// Read a file and return its contents as a slice of strings.
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fileString := string(fileBytes)
	return strings.Split(fileString, "\n"), nil
}

func StringsToGrid(input []string) map[int]map[int]string {
	// Convert a slice of strings to a 2D grid.
	grid := map[int]map[int]string{}
	for y, line := range input {
		grid[y] = map[int]string{}
		for x, char := range line {
			grid[y][x] = string(char)
		}
	}
	return grid
}
