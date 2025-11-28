// Escreva um algoritmo que simule o jogo da adivinhação

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	segredo := rand.Intn(100) + 1
	var palpite int
	for {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&palpite)
		if palpite < segredo {
			fmt.Println("Maior!")
		} else if palpite > segredo {
			fmt.Println("Menor!")
		} else {
			fmt.Println("Acertou!")
			break
		}
	}
}
