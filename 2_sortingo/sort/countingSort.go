package sort

import (
	U "../utils"
	"unsafe"
)

func max(input []int) int {
	if len(input) == 0 {
		return 0
	}
	max := input[0]
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}

func CountingSort(arr []int) {
	k := max(arr) + 1
	c := make([]int, k)

	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}
	for i, sum := 0, 0; i < k; i++ {
		sum, c[i] = sum+c[i], sum
	}
	sorted := make([]int, len(arr))
	for _, n := range arr {
		sorted[c[n]] = n
		c[n]++
	}
	copy(arr, sorted)
	U.Swaps = 0
	U.Comparisons = 0
	U.AddlMemory = (unsafe.Sizeof(c) + uintptr(k)*unsafe.Sizeof(k)) + (unsafe.Sizeof(sorted) + uintptr(len(arr))*unsafe.Sizeof(k))
}
