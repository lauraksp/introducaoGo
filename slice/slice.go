package main

import "fmt"

func main() {
	fmt.Println("Ola")
	var nomes = []string{"Laura", "Kimberly", "Joao"}

	for i, v := range nomes {
		fmt.Printf("ola %s voce está na posicao numero: %d\n", v, i)
	}
}
