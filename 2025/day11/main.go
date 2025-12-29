package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/basil127/advent-of-code/input"
)

const day string = "11"

const testFile string = "2025/day" + day + "/test.txt"
const inputFile string = "2025/day" + day + "/input.txt"

var cache map[string]int = make(map[string]int)

func parseGraph(lines []string) (map[string]struct{}, map[string][]string) {
	nodes := make(map[string]struct{})
	edges := make(map[string][]string)

	for _, line := range lines {
		from := line[:3]
		var to []string
		to = strings.Split(line[5:], " ")
		// add nodes
		nodes[from] = struct{}{}
		for _, t := range to {
			nodes[t] = struct{}{}
			// add edges
			edges[from] = append(edges[from], t)
		}

	}
	return nodes, edges
}

/*
BFS to find all paths from start to end
*/
func bfsPaths(start, end string, nodes map[string]struct{}, edges map[string][]string) int {
	// Iterative BFS enumerating simple paths using a queue of states.
	type state struct {
		node    string
		visited map[string]bool
	}

	count := 0
	// initialize queue with start node marked visited
	startVisited := make(map[string]bool)
	startVisited[start] = true
	queue := []state{{node: start, visited: startVisited}}

	for len(queue) > 0 {
		// pop front
		s := queue[0]
		queue = queue[1:]

		if s.node == end {
			count++
			continue
		}

		for _, neighbor := range edges[s.node] {
			if !s.visited[neighbor] {
				// copy visited map for new path state
				newVisited := make(map[string]bool, len(s.visited)+1)
				for k, v := range s.visited {
					newVisited[k] = v
				}
				newVisited[neighbor] = true
				queue = append(queue, state{node: neighbor, visited: newVisited})
			}
		}
	}

	return count
}

/*
DFS to find all paths from start to end that visit 'fft' and 'dac'
*/
func dfsPaths(current, end string, fftSeen, dacSeen bool, nodes map[string]struct{}, edges map[string][]string) int {
	if val, ok := cache[fmt.Sprintf("%s|%t|%t", current, fftSeen, dacSeen)]; ok {
		return val
	}
	if current == end && fftSeen && dacSeen {
		return 1
	}

	count := 0
	for _, neighbor := range edges[current] {
		if neighbor == "fft" && !fftSeen {
			count += dfsPaths(neighbor, end, true, dacSeen, nodes, edges)
		} else if neighbor == "dac" && !dacSeen {
			count += dfsPaths(neighbor, end, fftSeen, true, nodes, edges)
		} else if neighbor != "fft" && neighbor != "dac" {
			count += dfsPaths(neighbor, end, fftSeen, dacSeen, nodes, edges)
		}
	}
	cache[fmt.Sprintf("%s|%t|%t", current, fftSeen, dacSeen)] = count
	return count
}

/*
Finds number fo paths from start to end in the given graph.
*/
func part1(start, end string, nodes map[string]struct{}, edges map[string][]string) int {
	return bfsPaths(start, end, nodes, edges)
}

func part2(start, end string, nodes map[string]struct{}, edges map[string][]string) int {
	var totalPaths int

	totalPaths += dfsPaths(start, end, false, false, nodes, edges)

	return totalPaths
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

	fmt.Printf("Hello, Advent of Code 2025, Day %s!\n", day)
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(*file)
	nodes, edges := parseGraph(lines)

	// 2. process parts
	part1Result = part1("you", "out", nodes, edges)
	fmt.Println("Part 1:\t", part1Result)
	part2Result = part2("svr", "out", nodes, edges)

	// 3. output final results
	fmt.Println("Part 2:\t", part2Result)
}
