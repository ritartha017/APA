package main

import (
	"testing"
)

var result int 

type fn func(int) int

func test(f fn, val int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of func to prevent
		// the compiler eliminating the function call.
		r = f(val)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
func BenchmarkFib1_1(b *testing.B){ test(fib1, 1, b) }
func BenchmarkFib1_10(b *testing.B){ test(fib1, 10, b) }
func BenchmarkFib1_40(b *testing.B){ test(fib1, 40, b) }
//func BenchmarkFib1_55(b *testing.B){ test(fib1, 55, b) }
//func BenchmarkFib1_100(b *testing.B){ test(fib1, 100, b) }
//func BenchmarkFib1_1000(b *testing.B){ test(fib1, 1000, b) }

func BenchmarkFib2_1(b *testing.B){ test(fib2, 1, b) }
func BenchmarkFib2_10(b *testing.B){ test(fib2, 10, b) }
func BenchmarkFib2_40(b *testing.B){ test(fib2, 40, b) }
func BenchmarkFib2_55(b *testing.B){ test(fib2, 55, b) }
func BenchmarkFib2_100(b *testing.B){ test(fib2, 100, b) }
func BenchmarkFib2_1000(b *testing.B){ test(fib2, 1000, b) }

func BenchmarkFib3_1(b *testing.B){ test(fib3, 1, b) }
func BenchmarkFib3_10(b *testing.B){ test(fib3, 10, b) }
func BenchmarkFib3_40(b *testing.B){ test(fib3, 40, b) }
func BenchmarkFib3_55(b *testing.B){ test(fib3, 55, b) }
func BenchmarkFib3_100(b *testing.B){ test(fib3, 100, b) }
func BenchmarkFib3_1000(b *testing.B){ test(fib3, 1000, b) }

