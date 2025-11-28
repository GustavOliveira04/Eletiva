package main

import "testing"

func BenchmarkEPrimo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ePrimo(11)
	}
}