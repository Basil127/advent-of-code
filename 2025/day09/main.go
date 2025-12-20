package main

import (
	"flag"
	"fmt"

	"github.com/basil127/advent-of-code/input"
)

const day = "09"
const testFile = "2025/day" + day + "/test.txt"
const inputFile = "2025/day" + day + "/input.txt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
Find size of grid needed to contain all points
*/
func findGridSize(points [][2]int) (int, int) {
	var maxX, maxY int
	for _, p := range points {
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}
	return maxX + 1, maxY + 1
}

/*
Create a grid with 0 no tile, 1 red tile, 2 green tile
*/
func createGrid(points [][2]int) [][]int {
	gridX, gridY := findGridSize(points)
	grid := make([][]int, gridY)

	// Fill grid with 0s
	for i := range grid {
		grid[i] = make([]int, gridX)
	}

	// Place tiles
	prev := points[len(points)-1]
	for _, p := range points {
		grid[p[1]][p[0]] = 1

		// Draw line from prev to p
		dx := p[0] - prev[0]
		dy := p[1] - prev[1]
		steps := abs(dx)
		if abs(dy) > steps {
			steps = abs(dy)
		}
		if steps > 0 {
			stepX := dx / steps
			stepY := dy / steps
			for s := 1; s < steps; s++ {
				x := prev[0] + s*stepX
				y := prev[1] + s*stepY
				grid[y][x] = 2
			}
		}
		prev = p
	}
	return grid
}

/*
createPoints reads a list of strings representing 2D points in the format "x,y"
*/
func createPoints(lines []string) [][2]int {
	points := make([][2]int, 0, len(lines))
	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			panic("failed to read line: " + line + "\nError:\n" + err.Error())
		}
		points = append(points, [2]int{x, y})
	}
	return points
}

func part1(points [][2]int) int {
	var largest int
	var dx, dy, cur int

	for i, p := range points {
		for _, q := range points[i+1:] {
			dx = abs(p[0]-q[0]) + 1
			dy = abs(p[1]-q[1]) + 1
			cur = dx * dy
			if cur > largest {
				largest = cur
			}
		}
	}

	return largest
}

func part2(points [][2]int) int {
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
	points := createPoints(lines)
	grid := createGrid(points)

	// Debug output
	for _, row := range grid {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println(lines)
	fmt.Println(points)

	// 2. process parts
	part1Result = part1(points)
	part2Result = part2(points)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
