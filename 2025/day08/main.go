package main

import (
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/basil127/advent-of-code/dsu"
	"github.com/basil127/advent-of-code/heap"
	"github.com/basil127/advent-of-code/input"
)

const day = "08"
const testFile = "2025/day" + day + "/test.txt"
const inputFile = "2025/day" + day + "/input.txt"
const topn = 3

var numConnections int

type JunctionBox struct {
	id       int
	position [3]int
}

type Pair struct {
	distance float64
	fromID   int
	toID     int
}

func buildJunctionBoxs(lines []string) []JunctionBox {
	list := make([]JunctionBox, 0, len(lines))
	for i, val := range lines {
		var x, y, z int
		_, err := fmt.Sscanf(val, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic("failed to read line: " + strconv.Itoa(i) + "\nError:\n" + err.Error())
		}
		list = append(list, JunctionBox{
			id:       i,
			position: [3]int{x, y, z},
		})
	}
	return list
}

/*
distance computes the Euclidean distance between two junction boxes.
*/
func distance(a, b JunctionBox) float64 {
	dist := math.Pow(float64(a.position[0]-b.position[0]), 2)
	dist += math.Pow(float64(a.position[1]-b.position[1]), 2)
	dist += math.Pow(float64(a.position[2]-b.position[2]), 2)
	return math.Sqrt(dist)
}

func computeAllPairDistances(boxes []JunctionBox) []Pair {
	distances := make([]Pair, 0, len(boxes)*(len(boxes)-1)/2)

	// 2. Compute all pairwise distances
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			dist := distance(boxes[i], boxes[j])
			distances = append(distances, Pair{
				distance: dist,
				fromID:   boxes[i].id,
				toID:     boxes[j].id,
			})
		}
	}
	return distances
}

func part1(boxes []JunctionBox, connectPairsCount, topn int) int {
	// 1. Setup Disjoint Set and Priority Queue
	disjointSet := dsu.NewDSU(boxes)
	pairDistanceFn := func(p1, p2 Pair) bool {
		return p1.distance < p2.distance
	}
	pq := heap.NewHeap(pairDistanceFn)
	distances := computeAllPairDistances(boxes)
	// 3. Add all distances to the priority queue
	pq.Heapify(distances)
	// fmt.Println(pq)

	var err error
	var pair Pair
	var b1, b2 JunctionBox
	// 4. Connect the closest pairs until we reach the desired number of connections
	for i := 0; i < connectPairsCount; i++ {
		// Find the next pair
		pair, err = pq.Pop()
		if err != nil {
			panic("failed to pop from priority queue: " + err.Error())
		}
		b1 = boxes[pair.fromID]
		b2 = boxes[pair.toID]
		_, err := disjointSet.Connected(b1, b2)
		if err != nil {
			panic("failed to check connection: " + err.Error())
		}

		// Connect the two boxes
		err = disjointSet.Union(b1, b2)
		if err != nil {
			panic("failed to union sets: " + err.Error())
		}
	}

	// 5. Get sizes of all connected components
	allSets := disjointSet.GetSets()
	sizes := make([]int, 0, len(allSets))
	for _, members := range allSets {
		sizes = append(sizes, len(members))
	}

	// fmt.Println("All components:", allSets)
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})

	// 6. Compute product of sizes of top N largest components
	product := 1
	for i := 0; i < topn && i < len(sizes); i++ {
		product *= sizes[i]
	}

	return product
}

func part2(boxes []JunctionBox) int {
	// 1. Setup Disjoint Set and Priority Queue
	disjointSet := dsu.NewDSU(boxes)
	pairDistanceFn := func(p1, p2 Pair) bool {
		return p1.distance < p2.distance
	}
	pq := heap.NewHeap(pairDistanceFn)
	distances := computeAllPairDistances(boxes)

	// 3. Add all distances to the priority queue
	pq.Heapify(distances)

	var err error
	var pair Pair
	var b1, b2 JunctionBox
	// 4. Connect the closest pairs until we have a single connected component
	connected := disjointSet.CountSets() == 1
	// Find the next pair that are not already connected
	for !connected {
		// Get Pair with smallest distance
		pair, err = pq.Pop()
		if err != nil {
			panic("failed to pop from priority queue: " + err.Error())
		}
		b1 = boxes[pair.fromID]
		b2 = boxes[pair.toID]

		// Connect the two boxes
		err = disjointSet.Union(b1, b2)
		if err != nil {
			panic("failed to union sets: " + err.Error())
		}

		// Check if all boxes are now connected
		connected = disjointSet.CountSets() == 1
	}

	fmt.Printf("Last 2 Boxes:\n\t%v\n\t%v\n", b1, b2)
	return b1.position[0] * b2.position[0]
}

func main() {
	// Parse command line flags
	file := flag.String("file", "input", "input file to use: 'input' or 'test'")
	flag.Parse()

	if *file == "input" {
		fmt.Println("Running input file")
		*file = inputFile
		numConnections = 1000
	} else {
		fmt.Println("Running Test File")
		*file = testFile
		numConnections = 10
	}
	// 0. setup

	fmt.Printf("Hello, Advent of Code 2025, Day %s!\n", day)
	var part1Result, part2Result int
	// 1. load input
	lines := input.LoadInput(*file)
	boxes := buildJunctionBoxs(lines)
	// fmt.Println("boxes:\n", boxes)

	// 2. process parts
	part1Result = part1(boxes, numConnections, topn)
	part2Result = part2(boxes)

	// 3. output final results
	fmt.Println("Part 1:\t", part1Result)
	fmt.Println("Part 2:\t", part2Result)
}
