// Automata para aceptar cadenas de numeros
package main

import (
	"fmt"
	"os"
	"unicode"
)

var cadena string
var posicion int
var Token_Entrada byte

func main() {
	cadena = "(18-(@r*(2-4)/7)+$g"
	Token_Entrada = PrimerToken()
	expresion()
}

func PrimerToken() byte {
	posicion = 1
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
		fmt.Printf("Match: %c \n", Token_Entrada)
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
	} else if is_cov(Token_Entrada) {
		cov()
	} else {
		numero()
	}
}

func cov() {
	if Token_Entrada == '@' {
		HacerMatch('@')
		identificador()
	} else if Token_Entrada == '$' {
		HacerMatch('$')
		identificador()
	} else {
		panic("Error")
	}
}

func identificador() {
	letra()
	identificador_prima()
}

func identificador_prima() {
	if is_letter(Token_Entrada) {
		letra()
		identificador_prima()
	} else {
		//epsilon
	}
}

func letra() {
	if is_letter(Token_Entrada) {
		HacerMatch(Token_Entrada)
	} else {
		panic("Error")
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

func is_cov(tokenEntrada byte) bool {
	if tokenEntrada == '@' || tokenEntrada == '$' {
		return true
	} else {
		return false
	}
}

func is_letter(tokenEntrada byte) bool {
	if unicode.IsLetter(rune(tokenEntrada)) {
		return true
	} else {
		return false
	}
}
