package main

import "fmt"

//Generic
type Number interface {
	int | float64
}

func soma[T Number](valor map[string]T) T {
	var sum T

	for _, v := range valor {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(soma(map[string]int{"Mateus": 10, "Eduardo": 20}))
}
