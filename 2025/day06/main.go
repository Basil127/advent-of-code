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

func mulSlice(ints []int) int {
	var total int = 1
	for _, v := range ints {
		total *= v
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

func parseInput2(lines []string) [][]int {
	cur_num := make([]byte, 0, len(lines))
	nums := [][]int{}
	nums = append(nums, []int{})
	currentQ := 0

	for col := range lines[0] {
		cur_num = cur_num[:0]
		for row := range lines[:len(lines)-1] {
			if lines[row][col] != ' ' {
				cur_num = append(cur_num, lines[row][col])
			}
		}
		if len(cur_num) == 0 {
			nums = append(nums, []int{})
			currentQ++
			continue
		}
		num, err := strconv.Atoi(string(cur_num))
		if err != nil {
			panic("error parsing int: " + string(cur_num))
		}
		nums[currentQ] = append(nums[currentQ], num)
	}

	return nums
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
	fmt.Println(totals)

	return sumSlice(totals)
}

func part2(lines [][]int, operations []string) int {
	totals := make([]int, 0, len(lines))

	for i, line := range lines {
		if operations[i] == "*" {
			totals = append(totals, mulSlice(line))
		} else if operations[i] == "+" {
			totals = append(totals, sumSlice(line))
		} else {
			panic("unknown operator: %s" + operations[i])
		}
	}

	return sumSlice(totals)
}

func main() {
	fmt.Println("Hello, Advent of Code 2025, Day 06!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(inputFile)
	inputs, operations := parseInput(lines)
	inputs2 := parseInput2(lines)
	// fmt.Println("lines part 1:\n", inputs, "\nlines part 2:\n", inputs2, "\noperations:\n", operations)

	// 2. process parts
	part1Result = part1(inputs, operations)
	part2Result = part2(inputs2, operations)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
