package main

import "fmt"

func main() {
	var n, n1 int

	fmt.Println("digite um numero")
	fmt.Scan(&n, &n1)

	if n >= 0 && n1 >= 0 {
		fmt.Println("o resultado da multiplicao e", n*n1)
	} else {
		fmt.Println("o resultado de soma e", n+n1)
	}

}
