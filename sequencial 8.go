package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o número de dias trabalhados por um funcionário e o valor da sua diária
	//e calcule o seu salário.

	var t, d int
	fmt.Println("Digite a quantidade de dias trabalhados.")
	fmt.Scan(&t)

	fmt.Println("Digite o valor da diária.")
	fmt.Scan(&d)

	s := t * d
	fmt.Println("O valor do salário é", s)

}
