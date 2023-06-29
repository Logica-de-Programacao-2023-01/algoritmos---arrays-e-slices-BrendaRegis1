package main

import "fmt"

func main() {
	// Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a
	//seguinte tabela:
	//até 9 anos: mirim
	//10 a 13 anos: infantil
	//14 a 17 anos: juvenil
	//maiores de 18 anos: adulto

	var x int
	fmt.Println("Digite sua idade.")
	fmt.Scan(&x)

	if x <= 9 {
		fmt.Println("Mirim")
	} else if x >= 10 && x <= 13 {
		fmt.Println("Infantil")
	} else if x >= 14 && x <= 17 {
		fmt.Println("Juvenil")
	} else if x > 18 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Digite a idade em anos.")
	}

}
