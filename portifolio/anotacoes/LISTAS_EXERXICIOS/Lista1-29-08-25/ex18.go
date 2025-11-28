// Escreva um algoritmo que receba três matrizes de inteiros sem sinal, de 8bits, grave no disco um arquivo jpeg dessa imagem

package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func gerarImagemJPEG(r, g, b [][]uint8, nomeArquivo string) {
	altura := len(r)
	largura := len(r[0])
	img := image.NewRGBA(image.Rect(0, 0, largura, altura))

	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			cor := color.RGBA{r[y][x], g[y][x], b[y][x], 255}
			img.Set(x, y, cor)
		}
	}
	arquivo, _ := os.Create(nomeArquivo)
	defer arquivo.Close()
	jpeg.Encode(arquivo, img, nil)
}

func main() {
	r := [][]uint8{{255, 0}, {0, 255}}
	g := [][]uint8{{0, 255}, {0, 255}}
	b := [][]uint8{{0, 0}, {255, 255}}

	gerarImagemJPEG(r, g, b, "saida.jpg")
}
