package main

import "fmt"

func main() {
	//Faça um algoritmo que calcule a média aritmética de quatro números informados pelo
	//usuário.

	var número1 float64
	fmt.Println("Informe número 1.")
	fmt.Scan(&número1)

	var número2 float64
	fmt.Println("Informe número 2.")
	fmt.Scan(&número2)

	var número3 float64
	fmt.Println("Informe número 3.")
	fmt.Scan(&número3)

	var número4 float64
	fmt.Println("Informe número 4.")
	fmt.Scan(&número4)

	fmt.Print("A média aritmética dos 4 números informados é", número1/4+número2/4+número3/4+número4/4)

}
