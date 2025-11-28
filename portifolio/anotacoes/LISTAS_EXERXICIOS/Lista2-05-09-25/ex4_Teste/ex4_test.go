package main

import "testing"

func BenchmarkVerificarNumero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		verificarNumero(8)
	}
}