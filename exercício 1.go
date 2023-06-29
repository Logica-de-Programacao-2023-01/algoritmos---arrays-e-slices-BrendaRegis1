package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o nome, a idade e o peso de uma pessoa e imprima esses
	//valores na tela.
	var nome string
	fmt.Println("Qual é o seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Olá", nome)

	var idade int
	fmt.Println("Qual é a sua idade?")
	fmt.Scan(&idade)
	fmt.Println("Sua idade é", idade, "anos")

	var peso float32
	fmt.Println("Qual é o seu peso?")
	fmt.Scan(&peso)
	fmt.Println("Seu peso é", peso, "kg")
}
