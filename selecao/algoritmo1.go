package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Println("digite o primeiro numero")
	fmt.Scan(&n1)
	fmt.Println("digite o primeiro numero")
	fmt.Scan(&n2)

	if n1 > n2 {
		fmt.Println("o maior numero e", n1)
	} else {
		fmt.Println("o maior numero e", n2)
	}
}
