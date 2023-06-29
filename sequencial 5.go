package main

import "fmt"

func main() {
	//Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.
	var i float64
	fmt.Println("Qual é a sua idade?")
	fmt.Scan(&i)

	ID := i * 365
	fmt.Println("Sua idade em dias é", ID)
}
