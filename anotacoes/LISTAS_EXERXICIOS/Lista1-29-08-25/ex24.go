// Escreva um algoritmo que calcule o MMC.

package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}

func main() {
	fmt.Println("MMC de 12 e 18:", mmc(12, 18))
}
