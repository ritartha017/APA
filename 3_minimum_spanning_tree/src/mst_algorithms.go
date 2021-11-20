package main

import (
	"github.com/algds/uf"
	"sort"
)

func KruskalMST(graph [][]int, V int) [][]int {
	var edges []WeightedEdge
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if i == j {
				edges = append(edges, Edge{i, j, graph[i][j]})
			}
			if i != j && graph[i][j] == 0 {
				continue
			}
			edges = append(edges, Edge{i, j, graph[i][j]})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight() < edges[j].Weight()
	})
	u := uf.New(V)
	result := make([][]int, V)
	for i := range result {
		result[i] = make([]int, V)
	}
	for _, e := range edges {
		if !u.Connected(e.From(), e.To()) {
			u.Union(e.From(), e.To())
			result[e.From()][e.To()] = e.Weight()
		}
	}
	return result
}

func PrimMST(edges [][]int, V int) [][]int {
	parent := make([]int, V)
	key := make([]int, 0, V)
	mstSet := make([]bool, 0, V)

	for i := 0; i < V; i++ {
		key = append(key, MaxInt)
		mstSet = append(mstSet, false)
	}

	key[0] = 0
	parent = append(parent, -1)

	for count := 0; count < V-1; count++ {
		u := minKey(key, mstSet, V)
		mstSet[u] = true

		for v := 0; v < V; v++ {
			if edges[u][v] != 0 && mstSet[v] == false && edges[u][v] < key[v] {
				parent[v] = u
				key[v] = edges[u][v]
			}
		}
	}
	result := make([][]int, V)
	for i := range result {
		result[i] = make([]int, V)
	}
	for i := 1; i < V; i++ {
		result[parent[i]][i] = edges[i][parent[i]]
	}
	return result
}

func minKey(key []int, mstSet []bool, V int) int {
	var min int = MaxInt
	var min_index int
	for v := 0; v < V; v++ {
		if mstSet[v] == false && key[v] < min {
			min = key[v]
			min_index = v
		}
	}
	return min_index
}
