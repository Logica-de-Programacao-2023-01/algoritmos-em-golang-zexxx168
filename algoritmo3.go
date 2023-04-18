package main

import "fmt"

func main() {
	var p, alt float64
	fmt.Println("qual e seu peso")
	fmt.Scan(&p)
	fmt.Println("qual e sua altura")
	fmt.Scan(&alt)

	imc := p / (alt * alt)

	fmt.Println("seu imc e ", imc)

}
