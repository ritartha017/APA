package sort

import (
	U "../utils"
	"unsafe"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
			U.Comparisons++
			U.Swaps++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	U.Swaps++
	return lo
}

func QuickSort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		QuickSort(a, lo, p-1)
		QuickSort(a, p+1, hi)
		U.AddlMemory += unsafe.Sizeof(lo)
	}
}
