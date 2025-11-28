// Escreva um algoritmo que calcule fatorial de um numero N qualquer

package main

import "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}

func main() {
	fmt.Println("Fatorial(5):", fatorial(5))
}
