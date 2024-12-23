package main

//https://adventofcode.com/2024/day/1

import (
	"aoc24_go/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func splitLeftRight(input []string) ([]int, []int) {
	// The input is a list of strings, each string is a line with two numbers separated by a whitespace.
	// The output should be two lists of integers, one for the left number and one for the right number.
	left, right := []int{}, []int{}
	for _, line := range input {
		parts := strings.Fields(line)
		left_part, _ := strconv.Atoi(parts[0])
		right_part, _ := strconv.Atoi(parts[1])
		left = append(left, left_part)
		right = append(right, right_part)
	}
	return left, right
}

func solvePart1(input []string) int {
	// result as float because of the Abs function
	result := 0.0
	// split the input into left and right parts
	left, right := splitLeftRight(input)
	// sort each part
	slices.Sort(left)
	slices.Sort(right)
	// len(left) == len(right)
	for i := 0; i < len(left); i++ {
		// convert left[i] and right[i] to float64 to use Abs
		result += math.Abs(float64(left[i] - right[i]))
	}
	return int(result)
}

func solvePart2(input []string) int {
	result := 0
	//split the input into left and right parts
	left, right := splitLeftRight(input)
	rightCounts := utils.CountOccurrences(right)
	for _, l := range left {
		mul, ok := rightCounts[l]
		// if the key is not found, the multiplier is 0
		if !ok {
			mul = 0
		}
		result += l * mul
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("01/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
