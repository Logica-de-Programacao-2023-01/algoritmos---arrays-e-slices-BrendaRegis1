package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o valor do dólar em reais e um valor em dólares, e converta
	//esse valor para reais.

	var cambio float64
	fmt.Println("Informe o valor do dólar. ")
	fmt.Scan(&cambio)

	var dolares float64
	fmt.Println("Qual é o valor em dólares a ser convertido?")
	fmt.Scan(&dolares)

	r := cambio * dolares

	fmt.Println("O valor convertido para reais é", r)

}
