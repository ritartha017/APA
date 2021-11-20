package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"strconv"
	"time"
)

func main2Dense() {
	var V int = 5
	var repetitions int = 4
	var timeKd = make(map[int]time.Duration)
	var timeKs = make(map[int]time.Duration)
	var timePd = make(map[int]time.Duration)
	var timePs = make(map[int]time.Duration)

	i := 0
	for i < repetitions {
		fmt.Println("\nⅤ", color.Colorize(color.Yellow, strconv.Itoa(V)))

		fmt.Println("⟿", " Dense")
		graphD1 := generateGraphDense(V)
		start1 := time.Now()
		KruskalMST(graphD1, V)
		duration1 := time.Since(start1)
		fmt.Println("Kruskal", duration1)
		timeKd[V] = duration1

		start2 := time.Now()
		PrimMST(graphD1, V)
		duration2 := time.Since(start2)
		fmt.Println("Prim", duration2)
		timePd[V] = duration2

		fmt.Println("⟿", " Sparse")
		graphD2 := generateGraphSparse(V)
		start3 := time.Now()
		KruskalMST(graphD1, V)
		duration3 := time.Since(start3)
		fmt.Println("Kruskal", duration3)
		timeKs[V] = duration3

		start4 := time.Now()
		PrimMST(graphD2, V)
		duration4 := time.Since(start4)
		fmt.Println("Prim", duration4)
		timePs[V] = duration4
		i++
		V *= 2
	}
	fmt.Println("^-^")
}

func main() {
	// main2Dense()
	var V int = 50

	graphD1 := generateGraphDense(V)
	pprint(graphD1)
	toDotFile(graphD1, "grafdense.dot")
	go openDot("grafdense.dot")

	// graphD2 := generateGraphSparse(V)
	// pprint(graphD2)
	// toDotFile(graphD2, "grafsparse.dot")
	// go openDot("grafsparse.dot")

	res1 := PrimMST(graphD1, V) // !!
	pprint(res1)
	toDotFile(res1, "resdense.dot")
	go openDot("resdense.dot")

	// res2 := PrimMST(graphD2, V) // !!
	// pprint(res2)
	// toDotFile(res2, "ressparse.dot")
	// go openDot("ressparse.dot")

	time.Sleep(time.Second)
	fmt.Println("^-^")

}
