package main

import "fmt"

func main() {
	//Faça um algoritmo que leia dois números inteiros e mostre o maior deles.

	var x, y int
	fmt.Println("Digite o primeiro número.")
	fmt.Scan(&x)

	fmt.Println("Digite o segundo número.")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("O primeiro número é maior que o segundo.")
	} else if y > x {
		fmt.Println("O segundo número é maior que o primeiro.")
	} else {
		fmt.Println("Os números são iguais.")
	}
}
