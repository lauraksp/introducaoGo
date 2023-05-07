package main

import (
	"fmt"
)

type Humano interface {
	Fala(texto string)
	Sente(sentimento string)
}

type Robo struct {
	Nome           string
	DataFabricacao string
}

func (r *Robo) Fala(texto string) {
	fmt.Println(texto)
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Fala(texto string) {
	fmt.Println(texto)
}

func (p *Pessoa) Sente(sentimento string) {
	fmt.Println("Estou sentindo ", sentimento)
}

func CadastrarCPF(h Humano) {
	//h.CPF = "15100521608"
	fmt.Println("CPF cadastrado")
}

func main() {

	p := &Pessoa{}
	//r := &Robo{}

	p.Fala("Olá")
	//r.Fala("Olá, sou um robo") - robo n fala

	CadastrarCPF(p)
	//CadastrarCPF(r) - robo nao tem cpf

}
