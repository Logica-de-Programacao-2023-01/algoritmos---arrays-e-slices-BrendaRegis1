package main

import "fmt"

func main() {
	//Faça um algoritmo que leia a altura e o sexo de uma pessoa e mostre se ela está abaixo, dentro ou acima
	//do peso ideal, utilizando as fórmulas do exercício 3 da lista de algoritmos sequenciais.

	var altura, sexo, peso float64
	fmt.Println("Informe sua altura.")
	fmt.Scan(&altura)

	fmt.Println("Informe seu peso.")
	fmt.Scan(&peso)

	fmt.Println("Informe seu sexo.")
	fmt.Scan(&sexo)

	imc := peso / (altura * altura)
	if imc < 18.5 {
		fmt.Println("Abaixo do peso.")
	} else if imc >= 18.6 && imc <= 24.9 {
		fmt.Println("Peso Ideal.")
	} else if imc >= 25.0 && imc <= 29.9 {
		fmt.Println("Levemente acima do peso.")
	} else if imc >= 30.0 && imc <= 34.9 {
		fmt.Println("Obesidade grau 1.")
	} else if imc >= 35.0 && imc <= 39.9 {
		fmt.Println("Obesidade grau 2.")
	} else if imc >= 40.0 {
		fmt.Println("Obesidade grau 3.")
	} else {
		fmt.Println("Erro.")
	}

}
