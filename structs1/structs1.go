package main

import "fmt"

// boa pratica, comentar o codigo inicialmente c nome do metodo

// Pessoa, definindo struct - cria layout
type Pessoa struct {
	Nome  string
	Idade int
}

// NewUser recebe string e recebe novo valor
func NewUser(nome string) Pessoa {

	p := Pessoa{
		Nome:  nome,
		Idade: 22,
	}

	return p
}

func main() {
	//slice de Pessoa
	var pessoas = []Pessoa{Pessoa{"Laura", 22},
		{"Kimberly", 20},
		{"Joao", 23},
		{"Paulo", 17},
	}

	for _, v := range pessoas { // "i" estava antes de v, mas ignoramos com _(blank identify)
		fmt.Printf("Pessoa: %s\t\tIdade: %d\n", v.Nome, v.Idade)
	}
}
