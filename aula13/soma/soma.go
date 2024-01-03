package soma

import "fmt"

type Value interface {
	int | float64
}

func Soma[T Value](val1, val2 *T) T {
	return *val1 + *val2
}

func AlterarValor[T Value](valor *T) {
	*valor = 20
}

func ImprimeSoma[T Value](resultado T) {
	fmt.Println(resultado)
}
