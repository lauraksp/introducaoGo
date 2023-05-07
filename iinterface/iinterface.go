package main

import (
	"fmt"
)

type Humano interface {
	Fala(texto string)
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Fala(texto string) {
	fmt.Println(texto)
}

func cadastrarCPF(h HUmano) {
	h.CPF = "15100521608"
	fmt.Println("CPF cadastrado")
}

func main() {

	p := Pessoa{}

	p.Fala("Olá")

}
