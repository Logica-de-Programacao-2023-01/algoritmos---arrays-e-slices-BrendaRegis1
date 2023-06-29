package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

	var x int
	fmt.Println("Digite um número")
	fmt.Scan(&x)

	y := 2

	if x%y == 0 {
		fmt.Println("x é par.")
	} else {
		fmt.Println("x é ímpar.")

	}

}
