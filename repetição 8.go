package main

import "fmt"

//Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.

func main() {

	var x int

	fmt.Print("Digite um número positivo inteiro.")
	fmt.Scan(&x)

	for i := 0; i <= x; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
	}
}
