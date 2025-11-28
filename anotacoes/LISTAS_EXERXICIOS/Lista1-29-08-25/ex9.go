// Algoritmo que retorna o valor e o endereço de uma variavel

package main

import "fmt"

func valorEndereco() {
	x := 10
	fmt.Println("Valor:", x, " | Endereço:", &x)
}

func main() {
	valorEndereco()
}
