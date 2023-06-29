package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.
	var número1 int
	fmt.Println("Digite o número 1.")
	fmt.Scan(&número1)

	var número2 int
	fmt.Println("Digite o número 2.")
	fmt.Scan(&número2)

	var número3 int
	fmt.Println("Digite o número 3.")
	fmt.Scan(&número3)

	soma := número1 + número2 + número3
	fmt.Println("A soma é", soma)

}
