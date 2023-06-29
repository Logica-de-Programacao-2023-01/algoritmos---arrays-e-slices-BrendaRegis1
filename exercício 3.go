package main

import "fmt"

func main() {
	//Faça um algoritmo que calcule o volume de uma caixa com base, altura e profundidade
	//informados pelo usuário.
	var base float64
	fmt.Print("Informe o valor da base da caixa.")
	fmt.Scan(&base)

	var altura float64
	fmt.Println("Informe o valor da altura.")
	fmt.Scan(&altura)

	var profundidade float64
	fmt.Println("Informe o valor da profundidade.")
	fmt.Scan(&profundidade)

	fmt.Println("O volume da caixa é", base*altura*profundidade)
}
