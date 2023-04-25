package main

import "fmt"

func main() {
	var kg float64

	fmt.Println("qual seu peso")
	fmt.Scan(&kg)

	libra := kg * 2.20462

	fmt.Println("seu peso em libras e", libra)

}
