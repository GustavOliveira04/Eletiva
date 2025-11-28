// Escreva um algoritmo que conte a quantidade de vogais e consoantes de um texto

package main

import (
	"fmt"
	"strings"
)

func contarVogaisConsoantes(texto string) (int, int) {
	texto = strings.ToLower(texto)
	vogais := "aeiou"
	v, c := 0, 0
	for _, ch := range texto {
		if ch >= 'a' && ch <= 'z' {
			if strings.ContainsRune(vogais, ch) {
				v++
			} else {
				c++
			}
		}
	}
	return v, c
}

func main() {
	v, c := contarVogaisConsoantes("Olá mundo")
	fmt.Println("Vogais:", v, "Consoantes:", c)
}
