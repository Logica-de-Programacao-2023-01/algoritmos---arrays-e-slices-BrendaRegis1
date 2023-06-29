package main

import "fmt"

//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	var numeros = [3]int{1, 2, 3}

	// inicia variável soma com 0

	soma := 0

	// percorre o array e soma seus valores

	for i := 0; i < len(numeros); i++ {
		soma = soma + numeros[i]

	}

	fmt.Printf("A soma dos valores do array é %d", soma)

}
