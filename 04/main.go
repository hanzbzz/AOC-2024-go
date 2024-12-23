package main

import (
	"aoc24_go/utils"
	"fmt"
)

var directions = [][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, 1},  // up-right
	{-1, -1}, // up-left
}

func checkForWord(x int, y int, dx int, dy int, word string, input []string) bool {
	for i := 0; i < len(word); i++ {
		if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
			return false
		}
		if input[x][y] != word[i] {
			return false
		}
		x += dx
		y += dy
	}
	return true
}

func checkForXMAS(grid []string) bool {
	// 3x3 grid
	topDown := string(grid[0][0]) + string(grid[1][1]) + string(grid[2][2])
	bottomUp := string(grid[2][0]) + string(grid[1][1]) + string(grid[0][2])
	return (topDown == "MAS" || topDown == "SAM") && (bottomUp == "MAS" || bottomUp == "SAM")
}

func solvePart1(input []string) int {
	result := 0
	word := "XMAS"
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			for _, direction := range directions {
				if checkForWord(x, y, direction[0], direction[1], word, input) {
					result++
				}
			}
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0
	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			grid := []string{
				input[y][x : x+3],
				input[y+1][x : x+3],
				input[y+2][x : x+3],
			}
			if checkForXMAS(grid) {
				result++
			}
		}
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("04/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
