package main

import "fmt"

var (
	empregado bool //booleano
)

func main() {

	fmt.Println("Ola mundo")

	empregado = false
	var salario float64 = 2.5
	//var salario float64 = 5.0
	//var salario float64 = 0.50

	if salario > 3 {
		fmt.Println("é maior que 3k")
	} else if salario > 1 {
		fmt.Println("está no caminho")
	} else {
		fmt.Println("continuar estudando")
	}

}
