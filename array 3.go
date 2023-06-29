package main

import "fmt"

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {

	var numeros = [4]float64{4.5, 6.7, 3.2, 8.9}

	//A variável "produto" será usada para armazenar o resultado do produto dos elementos do array.
	produto := 1.0

	for i := 0; i < len(numeros); i++ {
		produto = produto * numeros[i]
	}

	fmt.Printf("O produto do array é %.2f", produto)

}
