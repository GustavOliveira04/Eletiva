package main

import "fmt"

func verificarNumero(n int) {
	if n%2 == 0 {
		fmt.Print("Par ")
	} else {
		fmt.Print("Impar ")
	}
	if n >= 0 {
		fmt.Print("e Positivo")
	} else {
		fmt.Print("e Negativo")
	}
}

func main() {
	verificarNumero(8)
}
