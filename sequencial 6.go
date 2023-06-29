package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 15%.

	var s float64
	fmt.Println("Digite o salário atual.")
	fmt.Scan(&s)

	NS := s + (s * 0.15)
	fmt.Println("Seu novo salário com aumento de 15% é", NS)
}
