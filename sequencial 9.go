package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.

	var p float64
	fmt.Println("Qual é o preço do produto?")
	fmt.Scan(&p)

	d := p - (p * 0.10)
	fmt.Println("O valor do produto com desconto de 10% é", d)

}
