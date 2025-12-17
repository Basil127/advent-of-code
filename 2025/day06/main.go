package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/basil127/advent-of-code/input"
)

const inputFile = "2025/day06/input.txt"

func filterEmptyStrings(input []string) []string {
	var n int
	for _, str := range input {
		if str == "" || str == " " {
			continue
		} else {
			input[n] = str
			n++
		}
	}
	return input[:n]
}

func sumSlice(ints []int) int {
	var total int
	for _, v := range ints {
		total += v
	}
	return total
}

func parseInput(lines []string) ([][]int, []string) {
	newlines := [][]int{}

	for _, line := range lines[:len(lines)-1] {
		newLine := strings.Split(line, " ")
		newLine = filterEmptyStrings(newLine)
		intList := make([]int, 0, len(newLine))
		for i := 0; i < len(newLine); i++ {
			val, err := strconv.Atoi(newLine[i])
			if err != nil {
				panic("error parsing int: " + newLine[i])
			}
			intList = append(intList, val)
		}
		newlines = append(newlines, intList)
	}

	operations := strings.Split(lines[len(lines)-1], " ")
	operations = filterEmptyStrings(operations)

	return newlines, operations
}

func part1(lines [][]int, operations []string) int {
	totals := make([]int, 0, len(lines[0]))
	totals = append(totals, lines[0]...)

	for _, line := range lines[1:] {
		for i, item := range line {
			switch operations[i] {
			case "*":
				totals[i] *= item
			case "+":
				totals[i] += item
			default:
				panic("unknown operator: %s" + operations[i])
			}
		}
	}

	return sumSlice(totals)
}

func part2(lines [][]int, operations []string) int {
	return 0
}

func main() {
	fmt.Println("Hello, Advent of Code 2025, Day 06!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(inputFile)
	inputs, operations := parseInput(lines)
	// fmt.Println("lines:\n", inputs, "\noperations:\n", operations)

	// 2. process parts
	part1Result = part1(inputs, operations)
	part2Result = part2(inputs, operations)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
