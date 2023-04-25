package main

import "fmt"

func main() {
	var dia int

	fmt.Println("Digite um numero de 1 a 7")
	fmt.Scan(&dia)

	switch dia {
	case 1:
		fmt.Println("domingo")
	case 2:
		fmt.Println("segunda-feira")
	case 3:
		fmt.Println("terca-feira")
	case 4:
		fmt.Println("quarta-feira")
	case 5:
		fmt.Println("quinta-feira")
	case 6:
		fmt.Println("sexta-feira")
	case 7:
		fmt.Println("sabado")
	default:
		fmt.Println("error digite 1 ou 7")
	}
}
