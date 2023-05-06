package main

import "fmt"

func main() {
	fmt.Println("Ola")
	var nomes = []string{"Laura", "Kimberly", "Joao"}
	var idades = []int{18, 20, 22}

	for i, v := range nomes {
		fmt.Printf("ola %s voce está na posicao numero: %d, e sua idade é %d\n", v, i, idades[i])
	}
}
