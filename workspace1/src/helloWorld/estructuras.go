package main

import "fmt"

type Token struct {
	Simbolo     string
	Tipo        string
	Ubicaciones []Ubication
}
type Ubication struct {
	Linea     int
	CaractIni int
	CaractFin int
}

func main() {
	suma := &Token{"+", "operador", []Ubication{Ubication{99, 1, 2}, Ubication{2, 1, 2}}}
	suma.Ubicaciones = append(suma.Ubicaciones, Ubication{78, 1, 2})
	fmt.Println(suma.Ubicaciones[2].Linea)
}
