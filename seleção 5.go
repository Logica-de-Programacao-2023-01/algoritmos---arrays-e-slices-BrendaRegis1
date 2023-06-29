package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

	var x int
	fmt.Println("Digite um número.")
	fmt.Scan(&x)

	y := 3
	z := 5

	if x%y == 0 && x%z == 0 {
		fmt.Println("O número é múltiplo de 3 e 5 ao mesmo tempo.")
	} else {
		fmt.Println("O número não é múltiplo de 3 e 5 ao mesmo tempo.")
	}
}
