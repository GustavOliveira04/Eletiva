package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const arquivoPPM = "p3.ppm"

	file, err := os.Open(arquivoPPM)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	formato := strings.TrimSpace(scanner.Text())
	if formato != "P3" {
		fmt.Println("Esse formato não é suportado! (use P3)")
		return
	}

	var tamanhoLinha string
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" && !strings.HasPrefix(linha, "#") {
			tamanhoLinha = linha
			break
		}
	}

	tamanho := strings.Fields(tamanhoLinha)
	largura, _ := strconv.Atoi(tamanho[0])
	altura, _ := strconv.Atoi(tamanho[1])

	scanner.Scan()
	maxValor, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	fmt.Printf("Formato: %s | Largura: %d | Altura: %d | Máx: %d\n\n", formato, largura, altura, maxValor)

	var valores []int
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha == "" || strings.HasPrefix(linha, "#") {
			continue
		}
		parts := strings.Fields(linha)
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			valores = append(valores, num)
		}
	}

	fmt.Println("Imagem criada:\n")
	i := 0
	pixel := 1
	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			if i+2 < len(valores) {
				r, g, b := valores[i], valores[i+1], valores[i+2]
				fmt.Printf("\x1b[48;2;%d;%d;%dm  \x1b[0m", r, g, b)
				i += 3
				pixel++
			}
		}
		fmt.Println()
	}

	fmt.Println("\nValores RGB de cada pixel:\n")
	i = 0
	pixel = 1
	for i+2 < len(valores) {
		r, g, b := valores[i], valores[i+1], valores[i+2]
		fmt.Printf("Pixel %d = R:%d G:%d B:%d\n", pixel, r, g, b)
		i += 3
		pixel++
	}
}
