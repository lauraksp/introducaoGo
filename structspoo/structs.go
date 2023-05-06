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

	nome := "Laura"
	pessoa := NewUser(nome)
	pessoa1 := NewUser("Paulo")
	pessoa2 := NewUser("Joao")
	fmt.Println(nome)

	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa.Nome, pessoa.Idade)
	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa1.Nome, pessoa1.Idade)
	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa2.Nome, pessoa2.Idade)

}
