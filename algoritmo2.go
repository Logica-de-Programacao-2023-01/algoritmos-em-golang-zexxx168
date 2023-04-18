package main

import "fmt"

func main() {
	var n1 int

	fmt.Println("escreva o seu numero")
	fmt.Scan(&n1)

	m2 := n1 * 2
	m3 := n1 * 3
	m4 := n1 * 4

	fmt.Println("dobro desse numero e", m2)
	fmt.Println("tripo desse numero e", m3)
	fmt.Println("quadruplo desse numero e", m4)

}
