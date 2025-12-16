package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "day01/input.txt"

func PositiveMod(a, b int) int {
	return (a%b + b) % b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
Splits the input string into lines, handling both Unix and Windows line endings.
*/
func LoadInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		println("Error reading file:", err)
		return []string{}
	}
	// Normalize Windows and Unix line endings and trim trailing newline
	s := strings.ReplaceAll(string(data), "\r\n", "\n")
	s = strings.TrimRight(s, "\n")
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "\n")
}

func processTurns(turns []string, start_num int) []int {
	results := []int{start_num}
	for i, turn := range turns {
		dir := turn[0]
		steps, err := strconv.Atoi(turn[1:])
		if err != nil {
			println("Error in turn", i, "with value", turn[1:], ":", turn)
			println("Error parsing steps:", err)
			continue
		}
		// Apply rotation
		switch dir {
		case 'L':
			results = append(results, PositiveMod(results[i]-steps, 100))
		case 'R':
			results = append(results, PositiveMod(results[i]+steps, 100))
		}
	}
	return results
}

func processTurnsPart2(turns []string, start_num int) int {
	cur := start_num
	count := 0
	for i, turn := range turns {
		var dir_int int
		dir := turn[0]
		steps, err := strconv.Atoi(turn[1:])
		if err != nil {
			println("Error in turn", i, "with value", turn[1:], ":", turn)
			println("Error parsing steps:", err)
			continue
		}
		// fmt.Println("Current:", results[i], "Turn:", turn)
		if dir == 'L' {
			dir_int = -1
		} else {
			dir_int = 1
		}
		// apply rotation
		count += abs(steps) / 100
		rem := dir_int * (steps % 100)
		cur += rem
		if cur >= 100 || (cur <= 0 && cur != rem) {
			count++
		}
		cur = PositiveMod(cur, 100)
	}

	fmt.Println("Inputs:\t", turns[:min(100, len(turns))])
	// fmt.Println("Result:\t", results[:min(100, len(results))])

	return count
}

func Count(values []int, target int) int {
	count := 0
	for _, v := range values {
		if v == target {
			count++
		}
	}
	return count
}

func part1() {
	fmt.Println("Hello, Advent of Code 2025!")

	lines := LoadInput(inputFile)
	fmt.Println("Number of lines in input:", len(lines))
	// fmt.Println("First line:", lines[0])

	results := processTurns(lines, 50)
	// fmt.Println("Results:", results)

	zerosCount := Count(results, 0)
	fmt.Println("Count of 0s in results:", zerosCount)
}

func part2() {
	fmt.Println("Hello, Advent of Code 2025!")

	lines := LoadInput(inputFile)
	fmt.Println("Number of lines in input:", len(lines))
	// fmt.Println("First line:", lines[0])

	result := processTurnsPart2(lines, 50)
	// fmt.Println("Results:", results)

	// zerosCount := Count(results, 0)
	fmt.Println("Count of 0s in results:", result)
}

func main() {
	part1()
	part2()
}
