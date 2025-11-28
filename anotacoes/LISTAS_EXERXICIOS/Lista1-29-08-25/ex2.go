// Algoritmo que divida dois numeros

package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Erro: Divisão por zero")
		return 0
	}
	return a / b
}

func main() {
	fmt.Println("Divisão:", divide(50, 2))
}
