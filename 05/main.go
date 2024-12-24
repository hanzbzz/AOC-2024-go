package main

import (
	"aoc24_go/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getRestrictionsAndUpdates(input []string) (map[string][]string, []string) {
	// value is slice of all items that need to be before the key
	restrictions := map[string][]string{}
	var updates []string
	for i, line := range input {
		// end of restrictions
		if line == "" {
			updates = input[i+1:]
			break
		}
		parts := strings.Split(line, "|")
		// X|Y means that X needs to be before Y in the update
		restrictions[parts[0]] = append(restrictions[parts[0]], parts[1])
	}
	return restrictions, updates
}

func isValidUpdate(parts []string, restrictions map[string][]string) (bool, int, int) {
	// returns bool indicating if the update is valid
	// if an update is invalid, return the index of offending parts (for part 2)
	for i, item := range parts {
		for j := 0; j < i; j++ {
			if slices.Contains(restrictions[item], parts[j]) {
				return false, i, j
			}
		}
	}
	return true, -1, -1
}

func solvePart1(input []string) int {
	result := 0
	restrictions, updates := getRestrictionsAndUpdates(input)
	for _, update := range updates {
		parts := strings.Split(update, ",")
		valid, _, _ := isValidUpdate(parts, restrictions)
		if valid {
			val, _ := strconv.Atoi(parts[len(parts)/2])
			result += val
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0
	restrictions, updates := getRestrictionsAndUpdates(input)
	for _, update := range updates {
		parts := strings.Split(update, ",")
		valid, _, _ := isValidUpdate(parts, restrictions)
		// skip valid updates
		if valid {
			continue
		}
		for {
			valid, i, j := isValidUpdate(parts, restrictions)
			if valid {
				val, _ := strconv.Atoi(parts[len(parts)/2])
				result += val
				break
			} else {
				// swap the offending parts
				// fmt.Println("Swapping", parts[i], "and", parts[j])
				parts[i], parts[j] = parts[j], parts[i]

			}
		}

	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("05/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
