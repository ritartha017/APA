package main

import (
	"reflect"
	"testing"
)

var V int = 5
var testGraph = [][]int{{0, 0, 0, 3, 12}, {0, 0, 2, 5, 0}, {0, 2, 0, 3, 7}, {3, 5, 3, 0, 0}, {12, 0, 7, 0, 0}}
var mstK = [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 2, 0, 0, 7}, {3, 0, 3, 0, 0}, {0, 0, 0, 0, 0}}
var mstP = [][]int{{0, 0, 0, 3, 0}, {0, 0, 0, 0, 0}, {0, 2, 0, 0, 7}, {0, 0, 3, 0, 0}, {0, 0, 0, 0, 0}}

func TestKruskalMST(t *testing.T) {
	t.Parallel()
	if result := KruskalMST(testGraph, V); !reflect.DeepEqual(result, mstK) {
		t.Errorf("Expected %v, got %v", mstK, result)
	}
}

func BenchmarkKruskalMST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KruskalMST(testGraph, V)
	}
}

func TestPrimMST(t *testing.T) {
	t.Parallel()
	if result := PrimMST(testGraph, V); !reflect.DeepEqual(result, mstP) {
		t.Errorf("Expected %v, got %v", mstP, result)
	}
}

func BenchmarkPrimMST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimMST(testGraph, V)
	}
}
