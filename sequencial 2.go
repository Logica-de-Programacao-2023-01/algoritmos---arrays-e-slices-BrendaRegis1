package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo

	var x int
	fmt.Println("Digite o número.")
	fmt.Scan(&x)

	dobro := 2 * x
	fmt.Println("O dobro é", dobro)

	triplo := 3 * x
	fmt.Println("O triplo é", triplo)

	quadruplo := 4 * x
	fmt.Println("O quadruplo é", quadruplo)
}
