package main

import (
	"aoc24_go/utils"
	"fmt"
	"strconv"
	"strings"
)

func blink(stones map[int]int) map[int]int {
	next := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			next[1] += count // Rule 1
		} else if len(strconv.Itoa(stone))%2 == 0 {
			stoneStr := strconv.Itoa(stone)
			tmp1, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			tmp2, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			next[tmp1] += count // Rule 2
			next[tmp2] += count
		} else {
			next[stone*2024] += count // Rule 3
		}
	}
	return next
}

func solve(input []string, numBlinks int) int {
	// map of count of each stone
	stones := make(map[int]int)
	for _, s := range strings.Fields(input[0]) {
		num, _ := strconv.Atoi(s)
		stones[num]++
	}
	// Simulate blinks
	for b := 0; b < numBlinks; b++ {
		stones = blink(stones)
	}
	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func main() {
	input, err := utils.ReadFileLines("11/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solve(input, 25)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solve(input, 75)
	fmt.Println("Part 2:", resultPart2)
}
