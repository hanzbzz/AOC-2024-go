package main

import (
	"aoc24_go/utils"
	"fmt"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	tmp := strconv.Itoa(a) + strconv.Itoa(b)
	result, _ := strconv.Atoi(tmp)
	return result
}

func calculateEquation(operands []int, operations []func(int, int) int) int {
	result := 0
	for i := 0; i < len(operations); i++ {
		if i == 0 {
			result += utils.Apply(operations[i], operands[i], operands[i+1])
		} else {
			result = utils.Apply(operations[i], result, operands[i+1])
		}
	}
	return result
}

func solve(input []string, part1 bool) int {
	result := 0
	for _, line := range input {
		parts := strings.Split(line, ":")
		testValStr, operandsStr := parts[0], parts[1]
		testVal, err := strconv.Atoi(testValStr)
		if err != nil {
			panic("Failed converting result")
		}
		operandsSlice := strings.Fields(operandsStr)
		var operands []int
		for _, opRaw := range operandsSlice {
			op, err := strconv.Atoi(opRaw)
			if err != nil {
				panic("Failed converting operand")
			}
			operands = append(operands, op)
		}
		var operationsCombinations [][]func(int, int) int
		if part1 {
			operationsCombinations = utils.GenerateCombinations([]func(int, int) int{add, mul}, len(operands)-1)
		} else {
			operationsCombinations = utils.GenerateCombinations([]func(int, int) int{add, mul, concat}, len(operands)-1)
		}
		for _, combination := range operationsCombinations {
			equationResult := calculateEquation(operands, combination)
			if equationResult == testVal {
				result += testVal
				break
			}
		}
	}
	return result
}

func main() {
	input, err := utils.ReadFileLines("07/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solve(input, true)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solve(input, false)
	fmt.Println("Part 2:", resultPart2)
}
