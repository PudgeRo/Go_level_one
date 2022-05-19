package fibb

import (
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	for i:=0; i < b.N; i++ {
		_ = Fibonacci(15)
	}
} 

func BenchmarkFibonacciWithOptimization(b *testing.B) {
	for i:=0; i < b.N; i++ {
		_ = FibonacciWithOptimization(15, map[int]int{0:0, 1:1})
	}
} 