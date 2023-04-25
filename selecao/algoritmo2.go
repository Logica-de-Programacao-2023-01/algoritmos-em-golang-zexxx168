package main

import "fmt"

func main() {

	var n1, n2, n3 int

	fmt.Println("Digite 3 numeros:")
	fmt.Scan(&n1, &n2, &n3)

	if n1 < n2 && n1 < n3 {
		println("Menor numero é:", n1)
	} else if n2 < n1 && n2 < n3 {
		fmt.Println("Menor numero é:", n2)
	} else {
		fmt.Println("Menor numero é:", n3)
	}
}
