//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fitness(ind, alvo string) int {
	score := 0
	for i := 0; i < len(alvo); i++ {
		if ind[i] == alvo[i] {
			score++
		}
	}
	return score
}

func crossover(a, b string) string {
	ponto := rand.Intn(len(a))
	return a[:ponto] + b[ponto:]
}

func mutacao(ind string, taxa float64) string {
	chars := []rune(ind)
	for i := 0; i < len(chars); i++ {
		if rand.Float64() < taxa {
			chars[i] = 'a' + rune(rand.Intn(26))
		}
	}
	return string(chars)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	alvo := "golang"
	tamanho := 20
	populacao := []string{}

	// inicial
	for i := 0; i < tamanho; i++ {
		ind := ""
		for j := 0; j < len(alvo); j++ {
			ind += string('a' + rune(rand.Intn(26)))
		}
		populacao = append(populacao, ind)
	}

	geracao := 0
	for {
		geracao++
		melhor, score := populacao[0], fitness(populacao[0], alvo)
		for _, ind := range populacao {
			if fitness(ind, alvo) > score {
				melhor, score = ind, fitness(ind, alvo)
			}
		}
		fmt.Println("Geração:", geracao, "Melhor:", melhor, "Score:", score)

		if melhor == alvo {
			fmt.Println("Encontrado!")
			break
		}

		novaPop := []string{}
		for i := 0; i < tamanho; i++ {
			pai := populacao[rand.Intn(tamanho)]
			mae := populacao[rand.Intn(tamanho)]
			filho := crossover(pai, mae)
			filho = mutacao(filho, 0.1)
			novaPop = append(novaPop, filho)
		}
		populacao = novaPop
	}
}
