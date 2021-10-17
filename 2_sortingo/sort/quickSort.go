package sort

import (
	U "../utils"
)

func partition(array []int, start int, end int) int {
	pivot_index := start
	pivot := array[pivot_index]
	for start < end {
		for start < len(array) && array[start] <= pivot {
			start = start + 1
			U.Comparisons += 2
		}
		for array[end] > pivot {
			end = end - 1
			U.Comparisons++
		}
		if start < end {
			array[start], array[end] = array[end], array[start]
			U.Swaps++
			U.Comparisons++
		}
	}
	array[end], array[pivot_index] = array[pivot_index], array[end]
	U.Swaps++
	return end
}

func QuickSort(array []int, start int, end int) {
	if start < end {
		p := partition(array, start, end)
		QuickSort(array, start, p-1)
		QuickSort(array, p+1, end)
	}
}
