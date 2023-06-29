package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números inteiros e mostre o menor deles.

	var x, y, z int
	fmt.Println("Digite o primeiro número.")
	fmt.Scan(&x)

	fmt.Println("Digite o segundo número.")
	fmt.Scan(&y)

	fmt.Println("Digite o terceiro número.")
	fmt.Scan(&z)

	menor := x

	if y < x {
		menor = y
	}
	if z < x {
		menor = z
	}
	fmt.Println("O menor número é", menor)
}
