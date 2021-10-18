package sort

import (
	U "../utils"
)

func flip(arr []int, i int) {
	start := 0
	for start < i {
		arr[start], arr[i] = arr[i], arr[start]
		start++
		i--
		U.Swaps++
	}
}

func findMax(arr []int, n int) int {
	mi := 0
	for i := 0; i < n; i++ {
		if arr[i] > arr[mi] {
			mi = i
		}
	}
	return mi
}

func PancakeSort(arr []int, n int) {
	for i := n; i > 1; i-- {
		mi := findMax(arr, i)
		if mi != i-1 {
			flip(arr, mi)
			flip(arr, i-1)
		}
	}
}
