// Escreva um algoritmo que retorne a quantidade de ocorrencias de uma palavra no texto, palavra é uma variavel

package main

import (
	"fmt"
	"strings"
)

func contarOcorrencias(texto, palavra string) int {
	texto = strings.ToLower(texto)
	palavra = strings.ToLower(palavra)
	return strings.Count(texto, palavra)
}

func main() {
	fmt.Println("Ocorrências:", contarOcorrencias("O gato e o cachorro", "o"))
}
