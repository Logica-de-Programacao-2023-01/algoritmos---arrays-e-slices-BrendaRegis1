package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.

	var x int
	fmt.Println("Digite um número.")
	fmt.Scan(&x)

	s := x + 1
	a := x - 1

	fmt.Println("Seu sucessor é", s)
	fmt.Println("Seu antecessor é", a)
}
