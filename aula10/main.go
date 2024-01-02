package main

import "fmt"

type Carro struct {
	Modelo string
	Cor    string
	valor  float64
}

type FuncoesCarro interface {
	NewCarro()
}

func (c Carro) NewCarro() {
	c.Modelo = "corola"
	c.Cor = "Preto"
	c.valor = 100000.00

	fmt.Println(c)
}

func salvar(f FuncoesCarro) {
	f.NewCarro()
}

func main() {
	c := Carro{}
	c.NewCarro()

	/* No caso eu cosigo passar um struct do tipo carro
	porque como o metodo que está atachado a struct carro
	é indetico ao metodo da interface então acontece uma
	implentação no caso a struct carro está implementando
	FuncoesCarro por isso funcoes carro é um carro
	*/
	salvar(c)
}
