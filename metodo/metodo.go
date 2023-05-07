package main

import "fmt"

type Site struct {
	Criador string
	URL     string
}

func (s *Site) DefineCriador(Name string) {
	s.Criador = Name
	fmt.Println("Dados salvos")
}

func (s *Site) DefineURL(URL string) {
	s.URL = URL
	fmt.Println("Dados salvos")
}

func main() {

	site := Site{} //obj declarado como vazio

	site.DefineCriador("Laura")
	site.DefineURL("www.medium.com/@laurakspp")
	fmt.Printf("Site %s\nCriador: %s\n", site.Criador, site.URL)
}
