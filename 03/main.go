package main

import (
	"aoc24_go/utils"
	"fmt"
	"regexp"
	"strconv"
)

func solvePart1(input []string) int {
	result := 0
	r, _ := regexp.Compile(`mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)`)
	//fmt.Println(r.FindAllString(input[0], -1))
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result += x * y
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0
	r, _ := regexp.Compile(`mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)|do\(\)|don't\(\)`)
	disabled := false
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				disabled = false
			} else if match[0] == "don't()" {
				disabled = true
			} else {
				if !disabled {
					x, _ := strconv.Atoi(match[1])
					y, _ := strconv.Atoi(match[2])
					result += x * y
				}
			}
		}
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("03/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
