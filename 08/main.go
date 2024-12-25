package main

import (
	"aoc24_go/utils"
	"errors"
	"fmt"
)

func getAntennaPositions(grid map[int]map[int]string) map[string][][2]int {
	// map where key is antenna name and value is list of positions
	antennaPositions := make(map[string][][2]int)
	for y, row := range grid {
		for x, val := range row {
			if val != "." {
				antennaPositions[val] = append(antennaPositions[val], [2]int{x, y})
			}
		}
	}
	return antennaPositions
}

func isPositionUniqueAntinode(p1 [2]int, p2 [2]int, grid map[int]map[int]string, antinodes map[[2]int]bool, multiplier int) (bool, error) {
	// get positions of antennas
	x1, y1, x2, y2 := p1[0], p1[1], p2[0], p2[1]
	dx, dy := (x1-x2)*multiplier, (y1-y2)*multiplier
	antinodeX, antinodeY := x1+dx, y1+dy
	// out of bounds
	_, ok2 := grid[antinodeY][antinodeX]
	if !ok2 {
		return false, errors.New("out of bounds")
	}
	// already marked as antinode
	if antinodes[[2]int{antinodeY, antinodeX}] {
		return false, nil
	}
	// new antinode
	antinodes[[2]int{antinodeY, antinodeX}] = true
	return true, nil
}

func solvePart1(input []string) int {
	result := 0
	grid := utils.StringsToGrid(input)
	antennaPositions := getAntennaPositions(grid)
	antinodePositions := make(map[[2]int]bool)
	for _, positions := range antennaPositions {
		for _, p1 := range positions {
			for _, p2 := range positions {
				// ignore identical antennas
				if p1 == p2 {
					continue
				}
				isAntinode, err := isPositionUniqueAntinode(p1, p2, grid, antinodePositions, 1)
				if err == nil && isAntinode {
					result++
				}
			}
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0
	grid := utils.StringsToGrid(input)
	antennaPositions := getAntennaPositions(grid)
	antinodePositions := make(map[[2]int]bool)
	for _, positions := range antennaPositions {
		for _, p1 := range positions {
			for _, p2 := range positions {
				// ignore identical antennas
				if p1 == p2 {
					continue
				}
				// distance multiplier
				// start with 0 to include the antenna itself
				multiplier := 0
				for {
					isAntinode, err := isPositionUniqueAntinode(p1, p2, grid, antinodePositions, multiplier)
					// out of bounds, no reason to continue
					if err != nil {
						break
					}
					if isAntinode {
						result++
					}
					multiplier++
				}
			}
		}
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("08/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
