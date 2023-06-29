package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC (Índice de Massa Corporal).
	var x, y float64
	fmt.Println("Informe seu peso.")
	fmt.Scan(&x)

	fmt.Println("Informe sua altura.")
	fmt.Scan(&y)

	IMC := x / (y * y)

	fmt.Println("Seu IMC é", IMC)
}
