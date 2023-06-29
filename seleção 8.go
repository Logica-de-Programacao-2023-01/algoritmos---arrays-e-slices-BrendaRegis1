package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da semana correspondente
	//(1 = domingo, 2 = segunda-feira, etc.).

	var x int
	fmt.Println("Digite um número entre 1 e 7:")
	fmt.Scan(&x)

	if x == 1 {
		fmt.Println("Domingo")
	} else if x == 2 {
		fmt.Println("Segunda-feira")
	} else if x == 3 {
		fmt.Println("Terça-feira")
	} else if x == 4 {
		fmt.Println("Quarta-feira")
	} else if x == 5 {
		fmt.Println("Quinta-feira")
	} else if x == 6 {
		fmt.Println("Sexta-feira")
	} else if x == 7 {
		fmt.Println("Sábado")
	} else {
		fmt.Println("Erro. Digite um número entre 1 e 7")
	}

}
