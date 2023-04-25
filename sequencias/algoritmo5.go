package main

import "fmt"

func main() {
	var idade int

	fmt.Println("qual sua idade ")
	fmt.Scan(&idade)

	anos := idade * 365

	fmt.Println("seus dias de vida sao", anos)
}
