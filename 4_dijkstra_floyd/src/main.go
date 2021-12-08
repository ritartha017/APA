package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"strconv"
	"time"
)

func main() {
	var V int = 1000
	src := 3
	graphDense := generateGraphDense(V)
	graphSparse := generateGraphSparse(V)
	var graphD, graphS [][]int
	sliceCopy(graphDense, &graphD)
	sliceCopy(graphSparse, &graphS)
	fmt.Println("\nⅤ", color.Colorize(color.Yellow, strconv.Itoa(V)))
	fmt.Println("⟿ Dense")
	{
		start1 := time.Now()
		cost, prev := DijkstraAllPaths(graphD, V, src)
		duration1 := time.Since(start1)
		fmt.Println("DIJKSTRA*", duration1)

		start2 := time.Now()
		dist, path := FloydWarshallSP(graphD, V)
		duration2 := time.Since(start2)
		fmt.Println("FLOYD   *", duration2)

		start3 := time.Now()
		costsrc, prevsrc := DijkstraSP(graphD, V, src)
		duration3 := time.Since(start3)
		fmt.Println("DIJKSTRA ", duration3)

		time.Sleep(time.Second) // !!!

		// print("\n[Dijkstra*]\n")
		// printSrcPathsDijkstra(prev, src, V)
		// time.Sleep(time.Second) // !!!
		// print("\n[Floyd*]\n")
		// printSrcPathsFloyd(path, src, V)
		// time.Sleep(time.Second) // !!!
		// print("\n[Dijkstra]\n")
		// printSrcPathsDijkstra(prevsrc, src, V)

		Use(cost, prev, dist, path, costsrc, prevsrc)
	}
	time.Sleep(time.Second) // !!!
	fmt.Print("\n⟿ Sparse")
	{
		start1 := time.Now()
		cost, prev := DijkstraAllPaths(graphS, V, src)
		duration1 := time.Since(start1)
		fmt.Println("\nDIJKSTRA*", duration1)

		start2 := time.Now()
		dist, path := FloydWarshallSP(graphS, V)
		duration2 := time.Since(start2)
		fmt.Println("FLOYD   *", duration2)

		start3 := time.Now()
		costsrc, prevsrc := DijkstraSP(graphS, V, src)
		duration3 := time.Since(start3)
		fmt.Println("DIJKSTRA ", duration3)

		time.Sleep(time.Second) // !!!

		Use(cost, prev, dist, path, costsrc, prevsrc)
	}
	// toDotFile(graphDense, "grafdense.dot")
	// go openDot("grafdense.dot")
	// toDotFile(graphSparse, "grafsparse.dot")
	// go openDot("grafsparse.dot")
	// time.Sleep(time.Second) // !!!
}
