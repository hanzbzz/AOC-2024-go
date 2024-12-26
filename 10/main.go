package main

import (
	"aoc24_go/utils"
	"fmt"
	"strconv"
)

var directions [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func scoreTrailHead(grid map[[2]int]int, trailHead [2]int, part1 bool) int {
	score := 0
	var helper func([2]int)
	// for part 1
	visited := make(map[[2]int]bool, 0)
	helper = func(position [2]int) {
		if grid[position] == 9 {
			if part1 {
				// only add to score if not yet visited
				if !visited[position] {
					score++
					visited[position] = true
				}
			} else {
				score++
			}
			return
		}
		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			neighborIdx := [2]int{position[0] + dy, position[1] + dx}
			if grid[neighborIdx] == grid[position]+1 {
				helper(neighborIdx)
			}
		}
	}
	helper(trailHead)
	return score
}

func solve(input []string, part1 bool) int {
	result := 0
	gridRaw := utils.StringsToGrid(input)
	grid := make(map[[2]int]int, 0)
	startingPositions := make([][2]int, 0)
	for y, row := range gridRaw {
		for x, col := range row {
			digit, _ := strconv.Atoi(col)
			grid[[2]int{y, x}] = digit
			// 0 -> trail head
			if digit == 0 {
				startingPositions = append(startingPositions, [2]int{y, x})
			}
		}
	}
	for _, trailHead := range startingPositions {
		result += scoreTrailHead(grid, trailHead, part1)
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("10/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solve(input, true)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solve(input, false)
	fmt.Println("Part 2:", resultPart2)
}
