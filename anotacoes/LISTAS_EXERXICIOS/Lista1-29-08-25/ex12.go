// Algoritmo que retorna se dois números são iguais ou não

package main

import "fmt"

func iguais(a, b int) bool {
	return a == b
}

func main() {
	fmt.Println("Iguais?", iguais(10, 11))
}
