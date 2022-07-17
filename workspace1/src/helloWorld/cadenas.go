package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "Hola mundo"
	//A mayusculas
	fmt.Println(strings.ToUpper(ss))
	//A minusculas
	fmt.Println(strings.ToLower(ss))

	//HasPrefix "si la cadena empieza con una serie de caracteres especificos"
	ss1 := "Hola"
	fmt.Println(strings.HasPrefix(ss, ss1))

	//HasSuffix "si la cadena termina con una serie de caracteres especificos"
	ss2 := "mundo"
	fmt.Println(strings.HasSuffix(ss, ss2))

	//Contains Busca si esta en la cadena
	ss3 := "do"
	fmt.Println(strings.Contains(ss, ss3))

	//Count Cuenta la candidad de veces que aparece en la cadena
	cad1 := "El maria el maria el el"
	fmt.Println(strings.Count(cad1, "El"))

	//Longitud de cadena, metodo len()
	cad2 := "El veloz murcielago hindu comia feliz cardillo y kiwi"
	fmt.Println(len(cad2))

	//Slipt, divison de car

}
