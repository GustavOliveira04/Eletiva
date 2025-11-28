// Algoritmo que faz ordenação em sequencia asc de caracteres

package main

import (
	"fmt"
	"sort"
	"strings"
)

func ordenaCaracteres(s string) {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	fmt.Println("Ordenado: ", strings.Join(chars, ""))
}

func main() {
	ordenaCaracteres("estou ordenado")
}
