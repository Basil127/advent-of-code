package main

import (
	"fmt"

	"github.com/basil127/advent-of-code/input"
)

const inputFile = "day04/input.txt"
const accessibleThreshold = 4

func constructGrid(lines []string) [][]int8 {
	var lineList []int8

	grid := make([][]int8, len(lines))
	for i, line := range lines {
		lineList = []int8{}
		var curInt int8
		for _, char := range line {
			if char == '.' {
				curInt = 0
			} else if char == '@' {
				curInt = 1
			}
			lineList = append(lineList, curInt)
		}
		grid[i] = lineList
	}
	return grid
}

func surroundingSum(grid [][]int8, row, col int) int8 {
	sum := int8(0)
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1}}
	for _, dir := range directions {
		drow := row + dir[0]
		dcol := col + dir[1]
		if drow >= 0 && drow < len(grid) && dcol >= 0 && dcol < len(grid[0]) {
			sum += grid[drow][dcol]
		}
	}
	return sum
}

func part2(grid [][]int8) int {
	accessible := 0
	new_accessible := 1
	for new_accessible > 0 {
		accessible += new_accessible
		new_accessible = 0
		for row := range grid {
			for col := range grid[row] {
				if grid[row][col] == 1 && surroundingSum(grid, row, col) < accessibleThreshold {
					new_accessible++
					grid[row][col] = 0
				}
			}
		}
	}
	return accessible - 1

	// 6491 too low, 7923 too high
}

func part1(grid [][]int8) int {
	accessilbe := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 && surroundingSum(grid, row, col) < accessibleThreshold {
				accessilbe++
			}
		}
	}
	return accessilbe
}

func main() {
	fmt.Println("Hello, Advent of Code 2025!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(inputFile)
	grid := constructGrid(lines)
	// fmt.Println("Grid:\n", grid)

	// 2. process parts
	part1Result = part1(grid)
	part2Result = part2(grid)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
