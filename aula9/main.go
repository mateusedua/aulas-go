package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
}

func (c Client) alteraIdade() {
	c.Idade = 20
	fmt.Println(c)
}

func main() {
	cliente := Client{Nome: "Mateus", Idade: 40}
	fmt.Println(cliente)
	cliente.alteraIdade()
}
