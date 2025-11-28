// Algoritmo que mostre o antecessor e o sucessor de um número

package main

import "fmt"

func sucessorAntecessor(n int) {
	fmt.Printf("Antecessor: %d | Sucessor: %d\n", n-1, n+1)
}

func main() {
	sucessorAntecessor(6)
}
