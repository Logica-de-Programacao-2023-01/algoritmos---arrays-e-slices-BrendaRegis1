package main

import "fmt"

func main() {
	//Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles,
	//se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

	var x, y int
	fmt.Println("Digite o primeiro número.")
	fmt.Scan(&x)

	fmt.Println("Digite o segundo número.")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("O resultado da multiplicação é", x*y)
	} else if x < 0 || y < 0 {
		fmt.Println("O resultado da soma é", x+y)
	}
}
