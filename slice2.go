package main

//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

import "fmt"

func main() {
	var tamanho int
	fmt.Println("Digite o tamanho do slice: ")
	fmt.Scan(&tamanho)
	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Println("Digite o valor do elemento: ", i+1)
		fmt.Scan(&slice[i])
	}

	fmt.Println("Slice: ", slice)

	soma := 0

	for _, valor := range slice {
		soma += valor
	}
	fmt.Println("A soma dos valores da slice é: ", soma)
}
