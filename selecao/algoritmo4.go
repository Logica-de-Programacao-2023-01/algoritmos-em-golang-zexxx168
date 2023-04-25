package main

import "fmt"

func main() {

	var altura float64
	var peso float64
	var sexo string

	fmt.Println("Qual sua altura?")
	fmt.Scan(&altura)
	fmt.Println("Qual seu peso?")
	fmt.Scan(&peso)
	fmt.Println("Qual seu sexo?(h ou m)")
	fmt.Scan(&sexo)

	imc := peso / (altura * altura)

	if imc < 18.5 && sexo == "h" {
		fmt.Println("Magreza")
	} else if imc >= 18.5 && imc <= 24.9 && sexo == "h" {
		fmt.Println("Peso normal")
	} else if imc >= 24.9 && imc <= 30 && sexo == "h" {
		fmt.Println("Sobrepeso")
	} else if imc < 18.5 && sexo == "m" {
		fmt.Println("Magreza")
	} else if imc >= 18.5 && imc <= 25 && sexo == "m" {
		fmt.Println("Peso normal")
	} else {
		fmt.Println("Sobrepeso")
	}

}
