package main

import (
	"fmt"

	"github.com/basil127/advent-of-code/input"
)

const inputFile = "2025/day00/input.txt"

func part1() int {
	return 0
}

func part2() int {
	return 0
}



func main() {
	fmt.Println("Hello, Advent of Code 2025, Day 00!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(inputFile)
	fmt.Println(lines)


	// 2. process parts
	part1Result = part1()
	part2Result = part2()

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}