package main

import (
	"aoc24_go/utils"
	"fmt"
	"strconv"
)

func mapToDisk(input string) ([]int, []int, []int) {
	// in: 2333133121414131402
	// free space will be -1
	// otherwise diskLayout[i] will be id of file block
	var diskLayout []int
	freeSpaceIndexes := make([]int, 0)
	fileBlockIndexes := make([]int, 0)
	for i, digitStr := range input {
		digit, err := strconv.Atoi(string(digitStr))
		if err != nil {
			panic("failed converting digit")
		}
		for j := 0; j < digit; j++ {
			if i%2 == 0 {
				diskLayout = append(diskLayout, i/2)
			} else {
				diskLayout = append(diskLayout, -1)
			}
		}
	}
	for i, digit := range diskLayout {
		if digit == -1 {
			freeSpaceIndexes = append(freeSpaceIndexes, i)

		} else {
			fileBlockIndexes = append(fileBlockIndexes, i)

		}
	}
	return diskLayout, fileBlockIndexes, freeSpaceIndexes
}

func moveFileBlock(disk []int, fileBlockIndex int, freeSpaceIndex int) []int {
	// simply swap disk[fileBlockIndex] with disk[freeSpaceIndex]
	if fileBlockIndex < freeSpaceIndex {
		return disk
	}
	lastFileBlockId := disk[fileBlockIndex]
	disk[freeSpaceIndex] = lastFileBlockId
	disk[fileBlockIndex] = -1
	return disk
}

func diskChecksum(disk []int) int {
	checksum := 0
	for i, c := range disk {
		if c == -1 {
			continue
		}
		checksum += i * c
	}
	return checksum
}

func solvePart1(input []string) int {
	disk, fileBlockIndexes, freeSpaceIndexes := mapToDisk(input[0])
	// i -> index of current freeSpaceIndexes
	// j -> index of current fileBlockIndexes
	// start free spaces from left and file blocks from right
	i, j := 0, len(fileBlockIndexes)-1
	for {

		// ran out of free spaces or file blocks, exit
		if i >= len(freeSpaceIndexes) || j >= len(fileBlockIndexes) {
			break
		}
		fileBlockIndex, freeSpaceIndex := fileBlockIndexes[j], freeSpaceIndexes[i]
		newDisk := moveFileBlock(disk, fileBlockIndex, freeSpaceIndex)
		i++
		j--
		disk = newDisk
	}
	return diskChecksum(disk)
}

func solvePart2(input []string) int {
	return 0
}

func main() {
	input, err := utils.ReadFileLines("09/input.txt")
	if err != nil {
		panic(err)
	}
	resultPart1 := solvePart1(input)
	fmt.Println("Part 1:", resultPart1)
	resultPart2 := solvePart2(input)
	fmt.Println("Part 2:", resultPart2)
}
