package main

import "fmt"

//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

func main() {
	inteiros := []int{1, 2, 3, 4, 5}

	inteiros = append(inteiros[:2], inteiros[3:]...)

	fmt.Println(inteiros)

}
