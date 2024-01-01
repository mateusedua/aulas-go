package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(2, 3))
	fmt.Println(teste(2, 3))

	// forma tambem de pegar os valores separados
	valor, error := retornaError("mateus")
	fmt.Println(valor, error)
}

func sum(a, b int) int {
	return a + b
}

// A função tambem pode retornar mais de um valor

func teste(a, b int) (int, bool) {
	if a > b {
		return a + b, true
	}

	return a - b, false
}

// Forma de retornar um erro

func retornaError(nome string) (bool, error) {
	if nome == "mateus" {
		return true, nil
	}

	return false, errors.New("O nome não é mateus")
}
