package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("qual o seu salario")
	fmt.Scan(&salario)

	if salario <= 1000.00 {
		salnovo := (10.0/100.0)*salario + salario

		fmt.Println("seu salario novo vai ser ", salnovo)
	} else {
		salnovo := (5.0/100.0)*salario + salario

		fmt.Println("seu salario novo vai ser ", salnovo)
	}

}
