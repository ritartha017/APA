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
}

var Swaps, Comparisons int
