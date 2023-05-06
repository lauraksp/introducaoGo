package main

import "fmt"

// boa pratica, comentar o codigo inicialmente c nome do metodo

// Teste imprime texto.
func Teste() {
	fmt.Println("Teste")
}

// NewUser recebe string e recebe novo valor
func NewUser(nome string) (string, int) {
	fmt.Println("Recebi: ", nome)

	return "Laura", 22
}

func main() {
	fmt.Println("Oii")

	nome := "Laura"
	pessoa, idade := NewUser(nome)
	fmt.Println(nome)
	fmt.Printf("Pessoa: %s\\t\tIdade: %d", pessoa, idade)
}
