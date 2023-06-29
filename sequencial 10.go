package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.
	var p float64
	fmt.Println("Qual é o seu peso em kg?")
	fmt.Scan(&p)

	l := p / 0.453592
	fmt.Println("Seu peso em libras é", l)
}
