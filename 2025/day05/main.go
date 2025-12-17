package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/basil127/advent-of-code/input"
)

const inputFile = "2025/day05/input.txt"

// interval represents an inclusive range [start,end]
type interval struct {
	start int
	end   int
}

/*
Return list of keys from map
*/
func getKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func parseInput(lines []string) (map[int]int, []int) {
	var validRanges = make(map[int]int)
	var ids []int
	for i, line := range lines {
		var err1, err2 error
		var id1, id2 int

		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			id1, err1 = strconv.Atoi(parts[0])
			id2, err2 = strconv.Atoi(parts[1])
			// If a start is repeated, keep the largest end, ranges can overlap and
			if prev, ok := validRanges[id1]; !ok || id2 > prev {
				validRanges[id1] = id2
			}
		} else if len(parts) == 1 && parts[0] != "" {
			id1, err1 = strconv.Atoi(parts[0])
			ids = append(ids, id1)
		}

		if err1 != nil || err2 != nil {
			panic(fmt.Sprintf("Error parsing line %d: \"%s\"", i+1, line))
		}

	}
	return validRanges, ids
}

/*
Return count of all fresh id's.
*/
func part2(ranges []interval) int {
	total := 0
	for _, iv := range ranges {
		total += (iv.end - iv.start + 1)
	}
	return total
}

// mergeRanges converts the map of ranges into a sorted, merged slice of
// non-overlapping intervals for fast lookup.
func mergeRanges(r map[int]int) []interval {
	keys := getKeys(r)
	sort.Ints(keys)
	ivs := make([]interval, 0, len(keys))
	for _, k := range keys {
		ivs = append(ivs, interval{start: k, end: r[k]})
	}

	if len(ivs) == 0 {
		return ivs
	}

	merged := make([]interval, 0, len(ivs))
	cur := ivs[0]
	for _, it := range ivs[1:] {
		if it.start <= cur.end+1 {
			// overlapping or adjacent — extend current
			if it.end > cur.end {
				cur.end = it.end
			}
		} else {
			merged = append(merged, cur)
			cur = it
		}
	}
	merged = append(merged, cur)
	return merged
}

func part1(ranges []interval, ids []int) int {
	sort.Ints(ids)

	validCount := 0
	for _, id := range ids {
		_, found := slices.BinarySearchFunc(ranges, id, func(iv interval, id int) int {
			if id < iv.start {
				return 1
			} else if id > iv.end {
				return -1
			} else {
				return 0
			}
		})

		if found {
			validCount++
		}
	}
	return validCount
}

func main() {
	fmt.Println("Hello, Advent of Code 2025!")
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(inputFile)
	validRanges, idList := parseInput(lines)
	ranges := mergeRanges(validRanges)

	// 2. process parts
	part1Result = part1(ranges, idList)
	part2Result = part2(ranges)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
