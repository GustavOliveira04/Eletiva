package main

import (
	"sort"
	"testing"
)

func BenchmarkOrdenaNumeros(b *testing.B) {
	sliceOriginal := []int{6, 8, 2, 1, 4, 0, 9, 3, 5, 7}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		nums := make([]int, len(sliceOriginal))
		copy(nums, sliceOriginal)
		sort.Ints(nums)
	}
}
