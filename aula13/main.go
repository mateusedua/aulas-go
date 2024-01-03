package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mateusedua/aulas-go/aula13/soma"
)

func main() {
	valor1 := 10
	valor2 := 10

	soma.AlterarValor(&valor1)
	result := soma.Soma(&valor1, &valor2)
	soma.ImprimeSoma(result)
	fmt.Println(uuid.New())
}
