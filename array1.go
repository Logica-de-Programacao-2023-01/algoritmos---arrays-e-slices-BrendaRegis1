package main

import "fmt"

//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.

func main() {
	soma := 0
	inteiros := [3]int{1, 2, 3}

	for _, valor := range inteiros {
		soma += valor

	}
	fmt.Println("A soma dos valores armazenados no array é:", soma)
}