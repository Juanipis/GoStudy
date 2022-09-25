// Automata para aceptar cadenas de numeros
package automatas

import (
	"fmt"
	"unicode"
)

var cadena string
var posicion int
var Token_Entrada byte

func main() {
	cadena1 := "12+((((($ab*67/2))*1))/8"
	cadena = cadena1

	posicion1 := 0
	posicion = posicion1

	var Token_Entrada1 byte = '0'
	Token_Entrada = Token_Entrada1

	Token_Entrada = PrimerToken()

	expresion()
	if Token_Entrada != ';' {
		if Token_Entrada == ')' {
			fmt.Println("Error: Hace falta '(' en algun lugar")
		} else {
			fmt.Printf("Error: Syntax: %c en la posición %d \n", Token_Entrada, posicion)
		}
		Token_Entrada = SiguienteToken()
		seguirExpresion()
	}

}

func seguirExpresion() {
	expresion_prima()
	if Token_Entrada == '*' || Token_Entrada == '/' {
		termino_prima()
	}
	if Token_Entrada != ';' {
		if Token_Entrada == ')' {
			fmt.Println("Error: Hace falta '(' en algun lugar")
		} else {
			fmt.Printf("Error: Syntax: %c en la posición %d \n", Token_Entrada, posicion)
		}
		Token_Entrada = SiguienteToken()
		seguirExpresion()
	}
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
		posicion = posicion + 1
		return ';'
	}
}

func HacerMatch(t byte) {
	if t == Token_Entrada {
		fmt.Printf("Match: %c %d\n", Token_Entrada, posicion)
		Token_Entrada = SiguienteToken()
	} else {
		fmt.Println("error")
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
		if Token_Entrada == ')' {
			HacerMatch(')')
		} else {
			fmt.Println("Error: Se esperaba un ')' en la posición:", posicion)
			SiguienteToken()
		}

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
		fmt.Println("Error: Se esperaba una variable o constante en la posición:", posicion)
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
		fmt.Println("Error: Se esperaba una letra en la posición:", posicion)
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
		fmt.Println("Error: Se esperaba un digito en la posición:", posicion)

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
