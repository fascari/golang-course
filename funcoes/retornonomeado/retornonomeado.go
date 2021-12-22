package main

import "fmt"

func trocar(p1, p2 int) (primeiro, segundo int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)

	primeiro, segundo := trocar(7, 1)
	fmt.Println(primeiro, segundo)
}
