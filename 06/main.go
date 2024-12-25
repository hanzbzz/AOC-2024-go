package main

import (
	"aoc24_go/utils"
	"errors"
	"fmt"
)

// each change is 90 degrees
// starting position is going up
// then go down, left, and up and repeat
var directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func findStartPosition(grid map[int]map[int]string) (int, int) {
	for y, row := range grid {
		for x, char := range row {
			if char == "^" {
				return x, y
			}
		}
	}
	return -1, -1
}

func getVisitedCount(visited map[[2]int]bool) int {
	result := 0
	for _, value := range visited {
		if value {
			result++
		}
	}
	return result
}

func performWalk(grid map[int]map[int]string, startX int, startY int) (map[[2]int]bool, error) {
	// true if grid[x][y] was visited
	visited := make(map[[2]int]bool)
	// direction with which the spot wat visited
	visitedDirection := make(map[[2]int]int)
	// number of turns made
	turns := 0
	currentX, currentY := startX, startY
	// don't move for the first iteration
	nextX, nextY := currentX, currentY
	for grid[nextY][nextX] != "" {
		// make the move
		currentX, currentY = nextX, nextY
		// check if we already visited this sport with same direction
		if visited[[2]int{currentY, currentX}] && visitedDirection[[2]int{currentY, currentX}]%len(directions) == turns%len(directions) {
			return nil, errors.New("found a cycle")
		}
		// mark current position as visited
		visited[[2]int{currentY, currentX}] = true
		visitedDirection[[2]int{currentY, currentX}] = turns
		// calculate next position with current direction
		nextX = currentX + directions[turns%len(directions)][0]
		nextY = currentY + directions[turns%len(directions)][1]
		// change directions until there is no obstacle
		for grid[nextY][nextX] == "#" {
			turns++
			nextX = currentX + directions[turns%len(directions)][0]
			nextY = currentY + directions[turns%len(directions)][1]
		}
	}
	return visited, nil
}

func solvePart1(input []string) int {
	grid := utils.StringsToGrid(input)
	startX, startY := findStartPosition(grid)
	visited, err := performWalk(grid, startX, startY)
	if err != nil {
		panic("Found cycle with no artificial obstacles")
	}
	return getVisitedCount(visited)
}

func solvePart2(input []string) int {
	result := 0
	grid := utils.StringsToGrid(input)
	startX, startY := findStartPosition(grid)
	// first perform walk with no additional obstacles
	visited, err := performWalk(grid, startX, startY)
	if err != nil {
		panic("Found cycle with no artificial obstacles")
	}
	// iterate over all visited positions
	for i := range visited {
		x, y := i[1], i[0]
		// can't place obstacle on starting position
		if grid[y][x] == "^" {
			continue
		}
		// place obstacle
		grid[y][x] = "#"
		_, err := performWalk(grid, startX, startY)
		// found cycle
		if err != nil {
			result++
		}
		// remove obstacle
		grid[y][x] = "."
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("06/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
