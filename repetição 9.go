package main

import "fmt"

//Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles.
//A leitura deve ser interrompida quando for lido o valor zero.

func main() {

	var x, soma, count int
	var media float64

	for {
		fmt.Println("Digite uma sequência de números inteiros: ")
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		// A cada iteração, o número digitado pelo usuário é adicionado à soma dos números e incrementa-se o
		//contador de números lidos.

		soma = soma + x
		count++
	}

	if count > 0 {
		media = float64(soma) / float64(count)
		fmt.Printf("A média aritmética dos números é %.2f\n ", media)

	} else {
		fmt.Println("Erro")
	}

}
