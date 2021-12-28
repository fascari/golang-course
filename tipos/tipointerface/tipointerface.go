package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

type dinamico interface{}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	printType(coisa)

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)
	printType(coisa2)

	coisa2 = true
	fmt.Println(coisa2)
	printType(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
	printType(coisa2)
}

func printType(coisa2 dinamico) {
	fmt.Println("Type:", reflect.TypeOf(coisa2))
}
