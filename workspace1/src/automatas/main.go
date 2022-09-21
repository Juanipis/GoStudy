// Automata para aceptar cadenas de numeros
package main

import (
	"fmt"
	"os"
)

var cadena string
var posicion int
var Token_Entrada byte

func main() {
	cadena = "1-(1*(2-4)/7)+3"
	posicion = 0
	Token_Entrada = PrimerToken()
	expresion()
}

func PrimerToken() byte {
	posicion = posicion + 1
	return cadena[0]
}

func SiguienteToken() byte {
	if posicion < len(cadena) {
		posicion = posicion + 1
		return cadena[posicion-1]
	} else {
		os.Exit(0)
		return 0
	}
}

func HacerMatch(t byte) {
	if t == Token_Entrada {
		fmt.Printf("Match: %c \t", Token_Entrada)
		Token_Entrada = SiguienteToken()
	} else {
		panic("Error")
	}
}

// Funciones de automata

func expresion() {
	termino()
	expresion_prima()
}

func expresion_prima() {
	if Token_Entrada == '+' {
		HacerMatch('+')
		termino()
		expresion_prima()
	} else if Token_Entrada == '-' {
		HacerMatch('-')
		termino()
		expresion_prima()
	} else {
		//epsilon
	}
}

func termino() {
	factor()
	termino_prima()
}

func termino_prima() {
	if Token_Entrada == '*' {
		HacerMatch('*')
		factor()
		termino_prima()
	} else if Token_Entrada == '/' {
		HacerMatch('/')
		factor()
		termino_prima()
	} else {
		//epsilon
	}
}

func factor() {
	if Token_Entrada == '(' {
		HacerMatch('(')
		expresion()
		HacerMatch(')')
	} else {
		numero()
	}
}

func numero() {
	digito()
	numero_prima()
}

func numero_prima() {
	if is_digit(Token_Entrada) {
		digito()
		numero_prima()
	} else {
		//epsilon
	}
}

func digito() {
	if is_digit(Token_Entrada) {
		HacerMatch(Token_Entrada)
	} else {
		panic("Error")
	}
}

func is_digit(tokenEntrada byte) bool {
	if tokenEntrada == '0' || tokenEntrada == '1' || tokenEntrada == '2' || tokenEntrada == '3' || tokenEntrada == '4' || tokenEntrada == '5' || tokenEntrada == '6' || tokenEntrada == '7' || tokenEntrada == '8' || tokenEntrada == '9' {
		return true
	} else {
		return false
	}
}
