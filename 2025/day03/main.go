package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/basil127/advent-of-code/input"
)

const inputFile = "day03/input.txt"

type MaxJoltage func(bank string) int

/*
Finds the maximum joltage using 2 digits from bank without changing order of occurence.
*/
func maxJoltage12(bank string) int {
	// 1. Convert to bank into list of ints
	numsList := []int{}
	for _, nStr := range bank {
		n, err := strconv.Atoi(string(nStr))

		if err != nil {
			println("Error parsing number:", string(nStr), ":", err)
			continue
		}
		numsList = append(numsList, n)
	}
	// joltages := []int{}
	total := 0
	start := 0
	for i := 11; i >= 0; i-- {
		cur_digit := slices.Max(numsList[start : len(numsList)-(i)])
		cur_digit_index := slices.Index(numsList[start:], cur_digit) + start
		start = cur_digit_index + 1
		// joltages = append(joltages, cur_digit)
		total += cur_digit * int(math.Pow(10, float64(i)))
	}

	return total
}

/*
Finds the maximum joltage using 2 digits from bank without changing order of occurence.
*/
func maxJoltage2(bank string) int {
	// 1. Convert to bank into list of ints
	numsList := []int{}
	for _, nStr := range bank {
		n, err := strconv.Atoi(string(nStr))
		if err != nil {
			println("Error parsing number:", string(nStr), ":", err)
			continue
		}
		numsList = append(numsList, n)
	}

	// 2. Greedy find first largest int, then second largest int after that
	firstDigit := slices.Max(numsList[:len(numsList)-1])
	firstDigitIndex := slices.Index(numsList, firstDigit)
	secondDigit := slices.Max(numsList[firstDigitIndex+1:])

	return (firstDigit * 10) + secondDigit
}

func allBatchMaxJoltages(banks []string, maxJoltage MaxJoltage) int {
	var joltage, totalJoltage int
	// joltages := []int{}

	for _, bank := range banks {
		joltage = maxJoltage(bank)
		fmt.Println("Bank:\t", bank, "Joltage:\t", joltage)
		// joltages = append(joltages, joltage)
		totalJoltage += joltage
	}
	return totalJoltage
}

func part1(banks []string) int {
	return allBatchMaxJoltages(banks, maxJoltage2)
}

func part2(banks []string) int {
	return allBatchMaxJoltages(banks, maxJoltage12)
}

func main() {
	var part1Result, part2Result int
	// 1. load input
	banks := input.LoadInput(inputFile)

	// 2. process parts
	part1Result = part1(banks)
	part2Result = part2(banks)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
