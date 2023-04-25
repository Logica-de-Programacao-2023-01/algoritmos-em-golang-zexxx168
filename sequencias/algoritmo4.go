package main

import "fmt"

func main() {
	var n1, n2, n3 float64

	fmt.Println("escreva o valor da sua nota")
	fmt.Scan(&n1)
	fmt.Println("escreva o valor da sua nota")
	fmt.Scan(&n2)
	fmt.Println("escreva o valor da sua nota")
	fmt.Scan(&n3)

	peso := n1 * 2
	peso2 := n2 * 3
	peso3 := n3 * 5
	peso4 := peso + peso2 + peso3
	peso5 := peso4 / 3

	fmt.Println("a sua media e", peso5)

}
