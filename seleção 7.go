package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10%
	//se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

	var x float64
	fmt.Println("Informe o salário.")
	fmt.Scan(&x)

	if x < 1000 {
		x *= 1.10
	} else {
		x *= 1.05
	}
	fmt.Println("O novo salário é", x)

}
