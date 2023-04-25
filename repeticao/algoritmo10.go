package main

import "fmt"

func main() {
	var n, m, maior int

	fmt.Println("digite o valor do numero")
	fmt.Scan(&n)
	for i := 0; i < 100; i++ {
		fmt.Println("digite outro sumero")
		fmt.Scan(&m)
		if m == 0 {
			break
		}
	}
	fmt.Println("o maior valor e")
}
