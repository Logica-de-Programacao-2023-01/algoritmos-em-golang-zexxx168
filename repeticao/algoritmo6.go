package main

import "fmt"

func main() {
	var numero int

	fmt.Println("digite um numero: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Println("2 x", i, "=", resultado)
	}
}
