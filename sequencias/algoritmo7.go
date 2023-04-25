package main

import "fmt"

func main() {
	var n1 int

	fmt.Println("digite um numero")
	fmt.Scan(&n1)

	sucessor := n1 + 1
	antecessor := n1 - 1

	fmt.Println("seu sucessor e", sucessor, "e seu antecessor", antecessor)
}
