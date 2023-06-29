package main

import "fmt"

// Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.

func main() {

	var x, y, z int

	fmt.Print("Digite um número: ")
	fmt.Scan(&x)

	fmt.Print("Digite um número: ")
	fmt.Scan(&y)

	fmt.Print("Digite um número: ")
	fmt.Scan(&z)

	if x > y && x > z {
		if y > z {
			fmt.Printf("A ordem é %d, %d, %d", z, y, x)
		} else {
			fmt.Printf("A ordem é %d, %d, %d", y, z, x)
		}
	} else if y > x && y > z {
		if x > z {
			fmt.Printf("A ordem é %d, %d, %d", z, x, y)
		} else {
			fmt.Printf("A ordem é %d, %d, %d", x, z, y)
		}
	} else if z > x && z > y {
		if x > y {
			fmt.Printf("A ordem é %d, %d, %d", y, x, z)
		} else {
			fmt.Printf("A ordem é %d, %d, %d", x, y, z)
		}
	}
}
