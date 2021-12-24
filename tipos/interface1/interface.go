package main

import (
	"fmt"
	"reflect"
)

type imprimivel interface {
	toString() string
	toStringTeste() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p produto) toStringTeste() string {
	return p.toString()
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p pessoa) toStringTeste() string {
	return p.toString()
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(reflect.TypeOf(x), "- toString() -", x.toString())
	fmt.Println(reflect.TypeOf(x), "- toStringTeste() -", x.toStringTeste())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	imprimir(p2)
}
