package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
## Task: **Problem Statement: Minimum Cost to Connect All Points**

**Description:** You are given a set of points on a 2D plane. Each point is represented as a pair of integers (x, y). The cost of connecting two points (x1, y1) and (x2, y2) is defined as the Manhattan distance between them: |x1 - x2| + |y1 - y2|.

Your task is to find the minimum cost to connect all points with each other, such that every point is connected to at least one other point. You can assume that there are no duplicate points.

**Input:** A list of points represented as pairs of integers [(x1, y1), (x2, y2), ..., (xn, yn)].

**Output:** Return the minimum cost to connect all the points.

**Constraints:**

- 1 <= points.length <= 1000
- 10^6 <= xi, yi <= 10^6
- All pairs (xi, yi) are distinct.

Example:

Input:
= [(0,0), (2,2), (3,10), (5,2), (7,0)]

Output: 20
Explanation:

Here's one possible way to connect the points with the minimum cost:

1. Connect (0,0) and (2,2) with a cost of 4 (|0-2| + |0-2|).
2. Connect (2,2) and (5,2) with a cost of 3 (|2-5| + |2-2|).
3. Connect (5,2) and (7,0) with a cost of 4 (|5-7| + |2-0|).
4. Connect (2,2) and (3,10) with a cost of 9 (|2-3| + |2-10|).

The total cost is 4 + 3 + 4 + 9 = 20.

Note that there may be other ways to connect the points with the same minimum cost.
*/

type Point struct {
	x, y int
}

type Edge struct {
	src, destination int
	cost             int
}

type MinHeap []*Edge

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Less(i, j int) bool { return m[i].cost < m[j].cost }

func (m *MinHeap) Pop() interface{} {
	temp := *m
	size := len(temp)
	min := temp[size-1]
	*m = temp[0 : size-1]
	return min
}

func (m *MinHeap) Push(i interface{}) {
	*m = append(*m, i.(*Edge))
}
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func calculateCost(i, j Point) int {
	return int(math.Abs(float64(i.x-j.x)) + math.Abs(float64(i.y-j.y)))
}

func calculateMinCostToConnectAllPoints(input []*Point) int {
	inputSize := len(input)
	if inputSize == 0 {
		return 0
	}

	mh := &MinHeap{}
	heap.Init(mh)
	visited := make([]bool, inputSize)
	visited[0] = true

	for i := 1; i < inputSize; i++ {
		cost := calculateCost(*input[0], *input[i])
		//add cost to min heap
		heap.Push(mh, &Edge{0, i, cost})
	}

	minimumCost := 0
	coveredEdges := 0
	//Prims algo
	for mh.Len() > 0 && (coveredEdges < inputSize) {
		edge := heap.Pop(mh).(*Edge)
		if visited[edge.destination] {
			continue
		}

		visited[edge.destination] = true
		minimumCost += edge.cost
		coveredEdges++

		//find all the edges from this point
		for i := 0; i < inputSize; i++ {
			if !visited[i] {
				dist := calculateCost(*input[edge.destination], *input[i])
				heap.Push(mh, &Edge{edge.destination, i, dist})
			}
		}
	}

	return minimumCost
}

func main() {
	fmt.Println("interview proven club")
	input := []*Point{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}

	result := calculateMinCostToConnectAllPoints(input)
	fmt.Printf("minimum cost to connect all points is %d", result)

}
