package main

import "fmt"

func main() {
	var n int

	fmt.Println("digite um numero")
	fmt.Scan(&n)

	fmt.Println("seus divisores de", n, "sao:")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			print(i, " ")
		}
	}
}
