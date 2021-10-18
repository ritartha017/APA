package utils

import (
	"time"
)

type Resources struct {
	SortingOrder string
	Size         int
	Swaps        int
	Comparisons  int
	Time         time.Duration
	AddlMemory   uintptr
}

var Swaps, Comparisons int
var AddlMemory uintptr
