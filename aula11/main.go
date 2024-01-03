package main

import "fmt"

type Cliente struct {
	Nome  string
	andar bool
}

//retornando uma strut ja apontada muito comum de se utilizar

func NewCliente() *Cliente {
	return &Cliente{Nome: "", andar: false}
}

func (c *Cliente) Andar() {
	c.andar = true
}

func main() {
	cliente := Cliente{Nome: "Mateus", andar: false}

	cliente.Andar()
	fmt.Println(cliente)

	teste := NewCliente()

	teste.Nome = "Mateus"
	teste.andar = false

	fmt.Println(teste)
}
