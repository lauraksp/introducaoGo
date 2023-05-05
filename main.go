package main

import "fmt"

var (
	empregado bool //booleano
)

func main() {
	fmt.Println("Ola mundo")

	nome := "Laura Kimberly"
	empregado = true

	if empregado {
		fmt.Println("Seja bem vinda, ", nome)
	} else {
		fmt.Println("Seja bem vinda")
	}
}
