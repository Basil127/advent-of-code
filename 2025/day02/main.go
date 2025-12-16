package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/basil127/advent-of-code/input"
)

const inputFile string = "day2/input.txt"

type ValidateCallable func(int) bool

/*
If the number is entirely a repeated sequence of digits (e.g., 1212 or 123123), return true.
*/
func containsRepeats(num int) bool {
	var repeatCandidate, fullRepeat string
	numStr := strconv.Itoa(num)
	length := len(numStr)
	for i := 1; i <= length/2; i++ {
		repeatCandidate = numStr[:i]
		fullRepeat = strings.Repeat(repeatCandidate, (length / i))
		if strings.Compare(fullRepeat, numStr) == 0 {
			return true
		}
	}
	return false
}

/*
Finds if the number is a double (e.g., 11, 22, 1212, 3333).
Is always false for odd length numbers.
*/
func isDouble(num int) bool {
	numStr := strconv.Itoa(num)
	if len(numStr)%2 == 0 && numStr[:len(numStr)/2] == numStr[len(numStr)/2:] {
		return true
	}
	return false
}

func sum(nums []string) int {
	sum := 0
	for _, n := range nums {
		val, err := strconv.Atoi(n)
		if err != nil {
			println("Error parsing number:", n, ":", err)
			continue
		}
		sum += val
	}
	return sum
}

func findInvalidIds(idRanges []string, validate ValidateCallable) []string {
	invalidIds := []string{}
	for i, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		if len(parts) != 2 {
			println("Invalid ID range format:", idRange, "indexed at", i)
			continue
		}
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			println("Error parsing ID range:", idRange, "indexed at", i)
			continue
		}
		for cur := start; cur <= end; cur++ {
			if validate(cur) {
				invalidIds = append(invalidIds, strconv.Itoa(cur))
			}
		}
	}
	return invalidIds
}

func part1() {
	fmt.Println("Hello, Advent of Code 2025 - Day 2!")

	lines := input.LoadInput(inputFile)
	idRanges := strings.Split(lines[0], ",")
	fmt.Println("Number of lines in input:", len(idRanges))
	// fmt.Println("First line:", lines[0])

	invalidIds := findInvalidIds(idRanges, isDouble)
	fmt.Println("Number of invalid IDs found:", len(invalidIds))
	fmt.Println("First 100 invalid IDs:", invalidIds[:min(100, len(invalidIds))])

	errorsTotal := sum(invalidIds)
	fmt.Println("Sum of invalid IDs:", errorsTotal)
}

func part2() {
	fmt.Println("Hello, Advent of Code 2025 - Day 2 Part 2!")

	lines := input.LoadInput(inputFile)
	idRanges := strings.Split(lines[0], ",")
	fmt.Println("Number of lines in input:", len(idRanges))
	// fmt.Println("First line:", lines[0])

	invalidIds := findInvalidIds(idRanges, containsRepeats)
	fmt.Println("Number of invalid IDs found:", len(invalidIds))
	fmt.Println("First 100 invalid IDs:", invalidIds[:min(100, len(invalidIds))])

	errorsTotal := sum(invalidIds)
	fmt.Println("Sum of invalid IDs:", errorsTotal)

	// less than 39988548038

}

func main() {
	part1()
	part2()
}
