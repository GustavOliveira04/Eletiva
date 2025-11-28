// Escreva um algoritmo que calcule o IMC, use como padrão uma pessoa brasileira.

package main

import "fmt"

func calcularIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func main() {
	fmt.Println("IMC:", calcularIMC(70, 1.75))
}
