package main

import (
	"aoc24_go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isSafe(parts []string) bool {
	// decided after first iteration
	increase, decrease := false, false
	//
	for i := 0; i < len(parts)-1; i++ {
		x, _ := strconv.Atoi(parts[i])
		y, _ := strconv.Atoi(parts[i+1])
		// first iteration, need to decide if increasing or decreasing
		if i == 0 {
			if x < y {
				increase = true
			} else if x > y {
				decrease = true
			}
		}
		// if increasing and the next number is smaller, return false
		if increase && x > y {
			return false
		}
		// if decreasing and the next number is bigger, return false
		if decrease && x < y {
			return false
		}
		// if the difference is not between [1,3], return false
		difference := int(math.Abs(float64(x - y)))
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func solve(input []string, part2 bool) int {
	result := 0
	for _, line := range input {
		parts := strings.Fields(line)
		if isSafe(parts) {
			result++
			continue
		}
		if part2 {
			// original line was not safe, try removing one number at a time
			for i := 0; i < len(parts); i++ {
				var partsCopy []string = make([]string, len(parts))
				copy(partsCopy, parts)
				partsCopy = utils.RemoveIndex(partsCopy, i)
				if isSafe(partsCopy) {
					result++
					break
				}
			}
		}
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("02/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solve(input, false)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solve(input, true)
	fmt.Println("Part 2:", resultPart2)
}
