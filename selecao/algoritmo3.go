package main

import "fmt"

func main() {
	var n1 int

	fmt.Println("digite um numero")
	fmt.Scan(&n1)

	if n1%2 == 0 {
		fmt.Println("esse numero e par", n1)
	} else {
		fmt.Println("esse numero e impar", n1)
	}
}
