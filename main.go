package main

import "fmt"

var (
	empregado bool //booleano
)

func main() {
	fmt.Println("Ola mundo")

	//nome := "Laura Kimberly"
	empregado = false
	var salario float64 = 2.500

	if salario > 3.000 {
		fmt.Println("é maior que 3k")
	} else if salario > 2.000 {
		fmt.Println("está no caminho")
	} else {
		fmt.Println("continuar estudando")
	}
}
