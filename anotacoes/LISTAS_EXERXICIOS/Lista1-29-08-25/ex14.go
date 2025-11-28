// Escreva um algoritmo que verifique se uma sequencia de caracteres forma um palindromo, funciona com frases tambem tirando os espaços

package main

import (
	"fmt"
	"strings"
)

func ehPalindromo(s string) bool {
	limpo := strings.ReplaceAll(strings.ToLower(s), " ", "")
	i, j := 0, len(limpo)-1
	for i < j {
		if limpo[i] != limpo[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println("É palíndromo?", ehPalindromo("A sacada da casa"))
}
