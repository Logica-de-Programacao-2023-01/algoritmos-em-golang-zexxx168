package main

import "fmt"

func main() {
	var n, s, s2 int
	fmt.Println("digite um sequencia de numeros")

	s = 0
	s2 = 0
	for i := 0; i < 100; i++ {
		fmt.Scan(&n)
		s = s + n
		s2 = s2 + 1
		if n == 0 {
			s2 = s2 - 1
			break
		}
	}
	m := s / s2
	fmt.Println("a media e", m)
}
