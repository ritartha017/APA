package main

import (
	S "./sort"
	U "./utils"
	"math/rand"
	"time"
)

func randomSlicegenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func increasingSlicegenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

func decreasingSlicegenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = size - i - 1
	}
	return slice
}

func performTesting(arr []int, n int, order string) U.Resources {
	var resources U.Resources
	U.Swaps, U.Comparisons = 0, 0

	start := time.Now()
	S.QuickSort(arr, 0, n-1)
	// S.MergeSort(arr)
	end := time.Since(start)

	resources.SortingOrder = order
	resources.Size = n
	resources.Time = end
	resources.Swaps = U.Swaps
	resources.Comparisons = U.Comparisons
	// unsafe.Sizeof(arr) + uintptr(n)*unsafe.Sizeof(n)
	return resources
}

func main() {
	var qSRandomSliceResources, qSIncreasingSliceResources, qSDecreasingSliceResources []U.Resources
	var table string

	var repetitions int = 4
	var iterations int = 0
	var n int = 10

	for iterations < repetitions {
		arr := make([]int, n, n)

		randomSlicegenerator(arr, n)
		// Resources taken by QS for random input
		res1 := performTesting(arr, n, "random")
		qSRandomSliceResources = append(qSRandomSliceResources, res1)

		increasingSlicegenerator(arr, n)
		// Resources taken by QS for increasing input
		res2 := performTesting(arr, n, "increasing")
		qSIncreasingSliceResources = append(qSIncreasingSliceResources, res2)

		decreasingSlicegenerator(arr, n)
		// Resources taken by QS for decreasing input
		res3 := performTesting(arr, n, "decreasing")
		qSDecreasingSliceResources = append(qSDecreasingSliceResources, res3)

		n *= 10
		iterations++
	}
	table = U.ToTable(qSRandomSliceResources, "QS - RandomSliceOutput")
	U.ToCSV(table, "RandomSliceOutput")

	table = U.ToTable(qSIncreasingSliceResources, "QS - IncreasingSliceOutput")
	U.ToCSV(table, "IncreasingSliceOutput")

	table = U.ToTable(qSDecreasingSliceResources, "QS - DecreasingSliceOutput")
	U.ToCSV(table, "DecreasingSliceOutput")

	U.DrawTimeChart(qSRandomSliceResources, qSIncreasingSliceResources, qSDecreasingSliceResources, "QS")
	U.DrawSwapsChart(qSRandomSliceResources, qSIncreasingSliceResources, qSDecreasingSliceResources, "QS")
	U.DrawComparisonsChart(qSRandomSliceResources, qSIncreasingSliceResources, qSDecreasingSliceResources, "QS")
}
