// Escreva um algoritmo que calcule a area de um retangulo

package main

import "fmt"

func areaRetangulo(base, altura float64) float64 {
	return base * altura
}

func main() {
	fmt.Println("Área:", areaRetangulo(5, 3))
}
