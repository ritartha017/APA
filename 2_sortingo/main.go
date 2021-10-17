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

func performTesting(arr []int, n int, order string, algorithm string) U.Resources {
	var resources U.Resources
	var start time.Time
	var end time.Duration
	U.Swaps, U.Comparisons = 0, 0

	if algorithm == "QS" {
		start = time.Now()
		S.QuickSort(arr, 0, n-1)
		end = time.Since(start)
	} else if algorithm == "MS" {
		start = time.Now()
		S.MergeSort(arr)
		end = time.Since(start)
	}
	resources.SortingOrder = order
	resources.Size = n
	resources.Time = end
	resources.Swaps = U.Swaps
	resources.Comparisons = U.Comparisons
	// unsafe.Sizeof(arr) + uintptr(n)*unsafe.Sizeof(n)
	return resources
}

func main() {
	var RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources []U.Resources
	var table string

	var repetitions int = 4
	var iterations int = 0
	var n int = 10

	var algName string = "QS"
	// -----------------------------------------------QS---------------------------------------------

	for iterations < repetitions {
		arr := make([]int, n, n)

		randomSlicegenerator(arr, n)
		// Resources taken by QS for random input
		res1 := performTesting(arr, n, "random", algName)
		RandomSliceResources = append(RandomSliceResources, res1)

		increasingSlicegenerator(arr, n)
		// Resources taken by QS for increasing input
		res2 := performTesting(arr, n, "increasing", algName)
		IncreasingSliceResources = append(IncreasingSliceResources, res2)

		decreasingSlicegenerator(arr, n)
		// Resources taken by QS for decreasing input
		res3 := performTesting(arr, n, "decreasing", algName)
		DecreasingSliceResources = append(DecreasingSliceResources, res3)

		n *= 10
		iterations++
	}
	table = U.ToTable(RandomSliceResources, algName+" - RandomSliceOutput")
	U.ToCSV(table, algName+"RandomSliceOutput")

	table = U.ToTable(IncreasingSliceResources, algName+" - IncreasingSliceOutput")
	U.ToCSV(table, algName+"IncreasingSliceOutput")

	table = U.ToTable(DecreasingSliceResources, algName+" - DecreasingSliceOutput")
	U.ToCSV(table, algName+"DecreasingSliceOutput")

	U.DrawTimeChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)
	U.DrawSwapsChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)
	U.DrawComparisonsChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)

	// -----------------------------------------------MS---------------------------------------------

	RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources = nil, nil, nil
	repetitions = 4
	iterations = 0
	n = 10

	algName = "MS"

	for iterations < repetitions {
		arr := make([]int, n, n)

		randomSlicegenerator(arr, n)
		res1 := performTesting(arr, n, "random", algName)
		RandomSliceResources = append(RandomSliceResources, res1)

		increasingSlicegenerator(arr, n)
		res2 := performTesting(arr, n, "increasing", algName)
		IncreasingSliceResources = append(IncreasingSliceResources, res2)

		decreasingSlicegenerator(arr, n)
		res3 := performTesting(arr, n, "decreasing", algName)
		DecreasingSliceResources = append(DecreasingSliceResources, res3)

		n *= 10
		iterations++
	}
	table = U.ToTable(RandomSliceResources, algName+" - RandomSliceOutput")
	U.ToCSV(table, algName+"RandomSliceOutput")

	table = U.ToTable(IncreasingSliceResources, algName+" - IncreasingSliceOutput")
	U.ToCSV(table, algName+"IncreasingSliceOutput")

	table = U.ToTable(DecreasingSliceResources, algName+" - DecreasingSliceOutput")
	U.ToCSV(table, algName+"DecreasingSliceOutput")

	U.DrawTimeChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)
	U.DrawSwapsChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)
	U.DrawComparisonsChart(RandomSliceResources, IncreasingSliceResources, DecreasingSliceResources, algName)
}
