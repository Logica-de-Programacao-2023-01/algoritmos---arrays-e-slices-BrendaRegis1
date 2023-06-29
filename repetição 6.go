package main

import "fmt"

func main() {
	//Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.

	var x int

	fmt.Println("Digite um número:")
	fmt.Scan(&x)

	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}

}
