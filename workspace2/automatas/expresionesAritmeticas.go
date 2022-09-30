// Automata para aceptar cadenas de numeros
package automatas

import (
	"strconv"
	"strings"
	"unicode"
)

// Variable: cadena
// Expresión aritmetica a evaluar
var cadena string

// Variable: posicion
// Posición en la cadena
var posicion int

// Variable: Token_Entrada
// Token de entrada
var Token_Entrada byte

// Variable: Log
// Log de errores
var Log string

/*
	 Function: AutomataExpresiones
		Automata para aceptar cadenas de expresiones aritmeticas del lenguaje messi

		Parameters:
			cadenaIN - Cadena de entrada a evaluar

		Returns:
			received - Booleano que indica si la cadena es aceptada o no
			Log - String con el log de errores

		See Also:
			<SiguienteToken>
			<PrimerToken>
			<expresion>
			<seguirExpresion>
*/
func AutomataExpresiones(cadenaIN string) (bool, string) {
	Log = ""
	cadena = cadenaIN
	Token_Entrada = PrimerToken()

	expresion()
	if Token_Entrada != ';' {
		if Token_Entrada == ')' {
			Log = Log + ("Error: Hace falta '(' en algun lugar\n")
		} else {
			Log = Log + ("Error: Syntax: " + string(Token_Entrada) + " en la posición " + strconv.Itoa(posicion) + "\n")
		}
		Token_Entrada = SiguienteToken()
		seguirExpresion()
	}

	var received bool
	if strings.Contains(Log, "Error") {
		received = false
	} else {
		received = true
	}

	return received, Log
}

/*
Function: seguirExpresion
*/
func seguirExpresion() {
	expresion_prima()
	if Token_Entrada == '*' || Token_Entrada == '/' || Token_Entrada == '%' || Token_Entrada == '^' {
		termino_prima()
	}
	if Token_Entrada != ';' {
		if Token_Entrada == ')' {
			Log = Log + ("Error: Hace falta '(' en algun lugar\n")
		} else {
			Log = Log + ("Error: Syntax: " + string(Token_Entrada) + " en la posición " + strconv.Itoa(posicion) + "\n")
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
		if cadena[posicion-1] == ' ' {
			return SiguienteToken()
		} else {
			return cadena[posicion-1]
		}
	} else {
		posicion = posicion + 1
		return ';'
	}
}

func HacerMatch(t byte) {
	if t == Token_Entrada {
		//Log = Log + ("Match: " + string(Token_Entrada) + " " + strconv.Itoa(posicion) + "\n")
		Token_Entrada = SiguienteToken()
	} else {
		Log = Log + ("error\n")
	}
}

func SiguienteTokenID() byte {
	if posicion < len(cadena) {
		posicion = posicion + 1
		return cadena[posicion-1]
	} else {
		posicion = posicion + 1
		return ';'
	}
}

func HacerMatchID(t byte) {
	if t == Token_Entrada {
		//Log = Log + ("Match: " + string(Token_Entrada) + " " + strconv.Itoa(posicion) + "\n")
		Token_Entrada = SiguienteTokenID()
	} else {
		Log = Log + ("error\n")
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
	} else if Token_Entrada == '%' {
		HacerMatch('%')
		factor()
		termino_prima()
	} else if Token_Entrada == '^' {
		HacerMatch('^')
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
		if unicode.IsLetter(rune(Token_Entrada)) {
			Log = Log + ("Error: Syntax: " + string(Token_Entrada) + " en la posición " + strconv.Itoa(posicion) + "\n")
			Token_Entrada = SiguienteToken()
			expresion_prima()
			if Token_Entrada == '*' || Token_Entrada == '/' || Token_Entrada == '%' || Token_Entrada == '^' {
				termino_prima()
			}
		}
		if Token_Entrada == ')' {
			HacerMatch(')')
		} else {
			Log = Log + ("Error: Se esperaba un ')' en la posición:" + strconv.Itoa(posicion) + "\n")
			Token_Entrada = SiguienteToken()
		}
	} else if is_cov(Token_Entrada) {
		if is_letter(cadena[posicion]) {
			cov()
		} else if is_digit(cadena[posicion]) {
			HacerMatch('@')
			numero()
		}
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
		Log = Log + ("Error: Se esperaba una variable o constante en la posición:" + strconv.Itoa(posicion) + "\n")
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
	} else if Token_Entrada == ' ' {
		Token_Entrada = SiguienteToken()
	} else {
		//epsilon
	}
}

func letra() {
	if is_letter(Token_Entrada) {
		HacerMatchID(Token_Entrada)
	} else {
		Log = Log + ("Error: Se esperaba una letra en la posición:" + strconv.Itoa(posicion) + "\n")
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
		Log = Log + ("Error: Se esperaba un digito, variable o constante en la posición:" + strconv.Itoa(posicion) + "\n")
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
