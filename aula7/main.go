package main

import "fmt"

func main() {
	salarios := map[string]int{"Marcos": 200, "Willian": 300, "Mateus": 500}

	// O make faz a inicialização zerada
	teste := make(map[string]int)
	// Outra forma de inicializar é não passar nada no valor
	teste2 := map[string]int{}
	//Usado para  deletar um valor do map
	delete(salarios, "Marcos")
	//Usado para incluir um valor no map
	salarios["Qualquer"] = 400

	fmt.Println(salarios, teste, teste2)
}
