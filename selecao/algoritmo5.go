package main

import "fmt"

func main() {
	var n1 int

	fmt.Println("digite um numero: ")
	fmt.Scan(&n1)

	if n1%3 == 0 && n1%5 == 0 {
		fmt.Println("esse numero e multiplo de 3 e de 5 ")
	} else {
		fmt.Println("esse numero n e multiplo de 3 e de 5 ")
	}
}
