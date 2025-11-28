package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func novaPessoa(nome string) *pessoa {
	p := pessoa{nome: nome}
	p.idade = 33
	return &p
}

func main() {
	fmt.Println(pessoa{"Gustavo", 21})

	fmt.Println(pessoa{nome: "Gabriel"})

	fmt.Println(novaPessoa("Fulano"))
}
