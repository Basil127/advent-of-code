package main

import (
	"flag"
	"fmt"

	"github.com/basil127/advent-of-code/input"
)

const day = "07"
const testFile = "2025/day" + day + "/test.txt"
const inputFile = "2025/day" + day + "/input.txt"

func part1(lines []string) int {
	window := []byte(lines[0])
	var splits int = 0

	for i, line := range lines[1:] {
		for j, char := range line {
			switch char {
			case '.':
				continue
			case '^':
				if window[j] == 'S' {
					splits++
					if j+1 < len(window) {
						window[j+1] = 'S'
					}
					if j-1 >= 0 {
						window[j-1] = 'S'
					}
					window[j] = '.'
				}
			default:
				panic("unknown char at line " + fmt.Sprint(i+1) + ", pos " + fmt.Sprint(j) + ": " + string(char))
			}
		}
	}

	return splits
}

func part2(lines []string) int {
	var splits int
	// 1. create sliding window
	window := make([]int, 0, len(lines[0]))
	for _, char := range lines[0] {
		switch char {
		case 'S':
			window = append(window, 1)
		case '.':
			window = append(window, 0)
		default:
			panic("unknown char in initial state: " + string(char))
		}
	}

	// 2. Go through each level
	for i, line := range lines[1:] {
		for j, char := range line {
			switch char {
			case '.':
				// No change in timelines
				continue
			case '^':
				// 3. Split timelines
				if window[j] > 0 {
					if j+1 < len(window) {
						window[j+1] += window[j]
					}
					if j-1 >= 0 {
						window[j-1] += window[j]
					}
					window[j] = 0
				}
			default:
				panic("unknown char at line " + fmt.Sprint(i+1) + ", pos " + fmt.Sprint(j) + ": " + string(char))
			}
		}
	}

	// 4. Sum splits
	for _, count := range window {
		splits += count
	}

	return splits
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
	// fmt.Println(lines)

	// 2. process parts
	part1Result = part1(lines)
	part2Result = part2(lines)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
