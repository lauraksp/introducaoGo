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

func main() {

	p := Pessoa{}

	p.Fala("Olá")

}
