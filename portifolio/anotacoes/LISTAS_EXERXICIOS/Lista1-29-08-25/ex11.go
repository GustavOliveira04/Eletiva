// Algoritmo que diz o dia da semana de acordo com o dia, mes e ano

package main

import (
	"fmt"
	"time"
)

func diaSemana(ano, mes, dia int) string {
	t := time.Date(ano, time.Month(mes), dia, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}

func main() {
	fmt.Println("Dia da semana:", diaSemana(2001, 11, 10)) // ANO, MES, DIA
}
