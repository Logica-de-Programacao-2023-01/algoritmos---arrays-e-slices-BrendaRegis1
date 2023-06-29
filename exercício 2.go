package main

import "fmt"

func main() {
	//Faça um algoritmo que calcule a área de um retângulo com base e altura informados
	//pelo usuário.
	var base float64
	fmt.Println("Qual é o valor da base do retângulo?")
	fmt.Scan(&base)
	fmt.Println("Sua base é", base)

	var altura float64
	fmt.Println("Qual é o valor da altura do retângulo?")
	fmt.Scan(&altura)
	fmt.Println("Sua altura é ", altura)

	fmt.Println("A área do retângulo é ", base*altura)
}
