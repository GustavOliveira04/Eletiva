// Escreva um algoritmo que faça conversao de unidades de temperatura

package main

import "fmt"

func celsiusParaFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}
func fahrenheitParaCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	fmt.Println("30°C =", celsiusParaFahrenheit(30), "°F")
	fmt.Println("86°F =", fahrenheitParaCelsius(86), "°C")
}
