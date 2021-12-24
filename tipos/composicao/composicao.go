package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	abrirCapota()
	// pode adicionar mais métodos
}

type bwm7 struct{}

func (b bwm7) abrirCapota() {
	fmt.Println("Abrindo capota...")
}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
	b.abrirCapota()
}
