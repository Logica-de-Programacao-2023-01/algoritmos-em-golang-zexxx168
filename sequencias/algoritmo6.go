package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("qual o seu salario")
	fmt.Scan(&salario)

	salnovo := (15.0/100.0)*salario + salario

	fmt.Println("seu salario novo vai ser ", salnovo)
}
