// Escreva um algoritmo que calcule a media entre dois ou mais numeros

package main

import "fmt"

func media(nums ...float64) float64 {
	soma := 0.0
	for _, n := range nums {
		soma += n
	}
	return soma / float64(len(nums))
}

func main() {
	fmt.Println("Média:", media(10, 20, 30))
}
