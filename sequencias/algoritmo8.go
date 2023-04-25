package main

import "fmt"

func main() {
	var dias, valordiario float64

	fmt.Println("dias trabalhados")
	fmt.Scan(&dias)
	fmt.Println("valordiario")
	fmt.Scan(&valordiario)

	salario := dias * valordiario

	fmt.Println("valor do salario", salario)

}
