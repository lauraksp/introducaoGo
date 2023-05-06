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

// FezAniversario recebe valor inteiro e atualiza ele
func FezAniverario(i *int) {
	*i = *i + 1
	fmt.Println("Recebi e atualizei", *i)
}

func main() {
	//slice de Pessoa
	var pessoas = []Pessoa{Pessoa{"Laura", 22},
		{"Kimberly", 20},
		{"Joao", 23},
		{"Paulo", 17},
	}

	a := pessoas[0].Nome
	i := pessoas[0].Idade

	FezAniverario(&i)

	fmt.Println("Valor resgatado: ", a, i)
}
