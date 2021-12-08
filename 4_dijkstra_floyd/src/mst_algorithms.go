package main

import (
	"fmt"
)

func DijkstraSP(graph [][]int, V int, src int) ([]int, []int) {
	prev := make([]int, V)
	cost := make([]int, 0, V)
	visited := make([]bool, 0, V)
	for i := 0; i < V; i++ {
		cost = append(cost, MaxInt)
		visited = append(visited, false)
	}

	cost[src] = 0
	prev[src] = -1

	for count := 1; count < V-1; count++ {
		u := minCost(cost, visited, V)
		visited[u] = true
		for v := 0; v < V; v++ {
			// if graph[u][v] != 0 && visited[v] == false && graph[u][v] < cost[v] {
			if !visited[v] && // graph[u][v] != 0 &&
				cost[u] != MaxInt && cost[u]+graph[u][v] < cost[v] {
				cost[v] = cost[u] + graph[u][v]
				prev[v] = u
				// cost[v] = graph[u][v]
			}
		}
	}
	return cost, prev
}

func FloydWarshallSP(graph [][]int, V int) ([][]int, [][]int) {
	dist := make([][]int, V)
	for i := 0; i < V; i++ {
		dist[i] = make([]int, V)
	}
	path := make([][]int, V)
	for i := 0; i < V; i++ {
		path[i] = make([]int, V)
	}
	for v := 0; v < V; v++ {
		for u := 0; u < V; u++ {
			if v != u && graph[v][u] == 0 {
				graph[v][u] = MaxInt
			}
		}
	}
	for v := 0; v < V; v++ {
		for u := 0; u < V; u++ {
			dist[v][u] = graph[v][u]
			if v == u {
				path[v][u] = 0
			} else if dist[v][u] != MaxInt {
				path[v][u] = v
			} else {
				path[v][u] = -1
			}
		}
	}
	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				if dist[i][k] != MaxInt && dist[k][j] != MaxInt &&
					dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					path[i][j] = path[k][j]
				}
			}
		}
	}
	return dist, path
}

func printPathDijkstra(parent []int, j int) {
	if parent[j] == -1 {
		return
	}
	printPathDijkstra(parent, parent[j])
	fmt.Print(j, " ")
}

func printPathFloyd(path [][]int, i, j int) {
	if path[i][j] == i {
		return
	}
	printPathFloyd(path, i, path[i][j])
	fmt.Print(path[i][j], " ")
}

func minCost(cost []int, mstSet []bool, V int) int {
	var min int = MaxInt
	var min_index int
	for v := 0; v < V; v++ {
		if mstSet[v] == false && cost[v] < min {
			min = cost[v]
			min_index = v
		}
	}
	return min_index
}

func DijkstraAllPaths(graphD [][]int, V, src int) ([]int, []int) {
	var cost, prev []int
	for i := 0; i < V; i++ {
		DijkstraSP(graphD, V, i)
		if i == src {
			cost, prev = DijkstraSP(graphD, V, i)
		}
	}
	return cost, prev
}

func printSrcPathsDijkstra(prev []int, src, V int) {
	for j := src; j < V; j++ {
		for i := 0; i < V; i++ {
			if i != src {
				fmt.Printf("[ %d ", j)
				printPathDijkstra(prev, i)
				fmt.Printf("]\n")
			}
		}
		if j == src {
			break
		}
	}
}

func printSrcPathsFloyd(path [][]int, src, V int) {
	for i := src; i < V; i++ {
		for j := 0; j < V; j++ {
			if i != j && path[i][j] != -1 {
				fmt.Printf("[ %d ", i)
				printPathFloyd(path, i, j)
				fmt.Printf("%d ]\n", j)
			}
		}
		if i == src {
			break
		}
	}
}
