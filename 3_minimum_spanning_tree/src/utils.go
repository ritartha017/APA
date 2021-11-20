package main

import (
	"fmt"
	"math/bits"
	"math/rand"
	"time"
)

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

const (
	MaxUint uint = (1 << bits.UintSize) - 1
	MaxInt  int  = (1<<bits.UintSize)/2 - 1
	MinInt  int  = (1 << bits.UintSize) / -2
)

func generateGraphDense(size int) [][]int {
	graph := make([][]int, size)
	for i := range graph {
		graph[i] = make([]int, size)
	}
	rand.Seed(time.Now().UnixNano())
	var max, min int = 10, 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				graph[i][j] = 0
				continue
			}
			graph[i][j] = rand.Intn(max-min) + min
			graph[j][i] = graph[i][j]
		}
	}
	return graph
}

func generateGraphSparse(size int) [][]int {
	graph := make([][]int, size)
	for i := range graph {
		graph[i] = make([]int, size)
	}
	rand.Seed(time.Now().UnixNano())
	var max, min int = 10, 1
	for i := 0; i < size; i++ {
		for j := 0; j < 2; j++ {
			graph[i][j] = rand.Intn(max-min) + min
			graph[j][i] = graph[i][j]
			graph[i][i] = 0
		}
	}
	return graph

}

func pprint(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("")
}
