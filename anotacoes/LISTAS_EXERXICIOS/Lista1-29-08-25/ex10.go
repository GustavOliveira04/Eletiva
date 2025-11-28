// Algoritmo que resolva a torre de hanoi com 3 discos

package main

import "fmt"

func hanoi(n int, origem, destino, auxiliar string) {
	if n == 1 {
		fmt.Printf("Mover disco de %s para %s\n", origem, destino)
		return
	}
	hanoi(n-1, origem, auxiliar, destino)
	fmt.Printf("Mover disco de %s para %s\n", origem, destino)
	hanoi(n-1, auxiliar, destino, origem)
}

func main() {
	hanoi(3, "A", "C", "B")
}
