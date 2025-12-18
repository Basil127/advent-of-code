package main

import (
	"flag"
	"fmt"

	"github.com/basil127/advent-of-code/input"
)

const day = "00"
const testFile = "2025/day" + day + "/test.txt"
const inputFile = "2025/day" + day + "/input.txt"

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func main() {
	// Parse command line flags
	file := flag.String("file", "input", "input file to use: 'input' or 'test'")
	flag.Parse()

	if *file == "input" {
		*file = inputFile
	} else {
		*file = testFile
	}
	// 0. setup

	fmt.Println("Hello, Advent of Code 2025, Day 00!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(*file)
	fmt.Println(lines)

	// 2. process parts
	part1Result = part1()
	part2Result = part2()

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
