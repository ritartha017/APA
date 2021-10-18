package main

import (
	S "./sort"
	U "./utils"
	"math/rand"
	"time"
)

func randomSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func positiveRandomSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}

func increasingSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

func decreasingSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = size - i - 1
	}
	return slice
}

func halfSortedSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size/2; i++ {
		slice[i] = i
	}
	for i := size / 2; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}

func almostSortedSliceGenerator(slice []int, size int) []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	// unsorted elem:
	slice[size/2] = 0
	return slice
}

func performTesting(arr []int, n int, order string, algorithm string) U.Resources {
	var resources U.Resources
	var start time.Time
	var end time.Duration
	U.Swaps, U.Comparisons, U.AddlMemory = 0, 0, 0

	if algorithm == "Quicksort" {
		start = time.Now()
		S.QuickSort(arr, 0, n-1)
		end = time.Since(start)
	} else if algorithm == "Mergesort" {
		start = time.Now()
		S.MergeSort(arr)
		end = time.Since(start)
	} else if algorithm == "Countingsort" {
		start = time.Now()
		S.CountingSort(arr)
		end = time.Since(start)
	} else if algorithm == "Pancakesort" {
		start = time.Now()
		S.PancakeSort(arr, len(arr))
		end = time.Since(start)
	}
	resources.SortingOrder = order
	resources.Size = n
	resources.Time = end
	resources.Swaps = U.Swaps
	resources.Comparisons = U.Comparisons
	resources.AddlMemory += U.AddlMemory
	return resources
}

func main() {
	var randomSliceResources, increasingSliceResources,
		decreasingSliceResources, halfSortedSliceResources,
		almostSortedSliceResources []U.Resources
	var table string

	var repetitions int = 4
	var iterations int = 0
	var n int = 10

	var algName string = "Pancakesort"

	// --------------------------------------------------------------------------------------------

	for iterations < repetitions {
		arr := make([]int, n, n)

		randomSliceGenerator(arr, n)
		// Resources taken by alg for random input
		res1 := performTesting(arr, n, "random", algName)
		randomSliceResources = append(randomSliceResources, res1)

		increasingSliceGenerator(arr, n)
		// Resources taken by alg for increasing input
		res2 := performTesting(arr, n, "increasing", algName)
		increasingSliceResources = append(increasingSliceResources, res2)

		decreasingSliceGenerator(arr, n)
		// Resources taken by alg for decreasing input
		res3 := performTesting(arr, n, "decreasing", algName)
		decreasingSliceResources = append(decreasingSliceResources, res3)

		halfSortedSliceGenerator(arr, n)
		// Resources taken by alg for half sorted input
		res4 := performTesting(arr, n, "half sorted", algName)
		halfSortedSliceResources = append(halfSortedSliceResources, res4)

		halfSortedSliceGenerator(arr, n)
		// Resources taken by alg for almost sorted input
		res5 := performTesting(arr, n, "almost sorted", algName)
		almostSortedSliceResources = append(almostSortedSliceResources, res5)

		n *= 10
		iterations++
	}
	table = U.ToTable(randomSliceResources, algName+" - RandomSliceOutput")
	U.ToCSV(table, algName+"RandomSliceOutput")

	table = U.ToTable(increasingSliceResources, algName+" - IncreasingSliceOutput")
	U.ToCSV(table, algName+"IncreasingSliceOutput")

	table = U.ToTable(decreasingSliceResources, algName+" - DecreasingSliceOutput")
	U.ToCSV(table, algName+"DecreasingSliceOutput")

	table = U.ToTable(halfSortedSliceResources, algName+" - HalfSortedSliceOutput")
	U.ToCSV(table, algName+"HalfSortedSliceOutput")

	table = U.ToTable(almostSortedSliceResources, algName+" - AlmostSortedSliceOutput")
	U.ToCSV(table, algName+"AlmostSortedSliceOutput")

	U.DrawTimeChart(randomSliceResources, increasingSliceResources,
		decreasingSliceResources, halfSortedSliceResources, almostSortedSliceResources, algName)
	U.DrawSwapsChart(randomSliceResources, increasingSliceResources,
		decreasingSliceResources, halfSortedSliceResources, almostSortedSliceResources, algName)
	U.DrawComparisonsChart(randomSliceResources, increasingSliceResources,
		decreasingSliceResources, halfSortedSliceResources, almostSortedSliceResources, algName)

}
