package main

import "fmt"

func main() {
	var produto float64

	fmt.Println("qual o valor do produto")
	fmt.Scan(&produto)

	valornovo := 0.1 * produto
	valornovo1 := produto - valornovo

	fmt.Println("valor aplicado do desconto e ", valornovo1)

}
