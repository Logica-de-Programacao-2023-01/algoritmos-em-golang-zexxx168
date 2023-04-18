package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Println("digite primeiro numero")
	fmt.Scan(&n1)
	fmt.Println("digite primeiro numero")
	fmt.Scan(&n2)
	fmt.Println("digite primeiro numero")
	fmt.Scan(&n3)

	soma := n1 + n2 + n3

	fmt.Println("sua soma e", soma)
}
