package main

import "fmt"

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.

func main() {

	array := [4]float64{1.2, 3.4, 5.5, 6.4}

	produto := 1.0

	for _, valor := range array {
		produto *= valor
	}
	fmt.Println(produto)
}
