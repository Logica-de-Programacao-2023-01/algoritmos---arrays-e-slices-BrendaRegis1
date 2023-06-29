package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles, com pesos 2, 3 e 5,
	//respectivamente.

	var x1, x2, x3 float64
	fmt.Println("Digite o primeiro número.")
	fmt.Scan(&x1)

	fmt.Println("Digite o segundo número.")
	fmt.Scan(&x2)

	fmt.Println("Digite o terceiro número.")
	fmt.Scan(&x3)

	MP := (x1*2)/10 + (x2*3)/10 + (x3*5)/10
	fmt.Println("A média ponderada entre esses números é", MP)
}
