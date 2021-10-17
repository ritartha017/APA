package sort

import (
	U "../utils"
)

func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
			U.Comparisons += 2
			U.Swaps++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
			U.Comparisons += 2
			U.Swaps++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
			U.Comparisons++
			U.Swaps++
		} else {
			slice[k] = right[j]
			j++
			U.Swaps++
		}
	}
	return slice
}
