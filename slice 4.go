package main

//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

import "fmt"

func main() {

	var tamanho int

	fmt.Println("Informe o tamanho do slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, 0, tamanho)

	for i := 0; i < tamanho; i++ {
		var elemento int
		fmt.Printf("Informe o valor do elemento %d: ", i+1)
		fmt.Scan(&elemento)

		slice = append(slice, elemento)
	}

	soma := 0
	for _, valor := range slice {
		soma = soma + valor
	}
	fmt.Printf("A soma do slice é %d", soma)
}
