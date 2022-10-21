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

var preFija string
var postFija string

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
func AutomataExpresiones(cadenaIN string) (bool, string, string, string) {
	Log = ""
	cadena = cadenaIN
	Token_Entrada = PrimerToken()
	preFija = ""
	postFija = ""

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

	return received, Log, preFija, postFija
}

/*
Function: seguirExpresion

	Función que sigue analizando en caso de que encuentre un error y se recupere. Se utilizan las variables globales cadena, posicion, log y Token_Entrada

	See Also:
		<expresion_prima>
		<termino_prima>
		<SiguienteToken>
		<seguirExpresion>
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

/*
	 Function: PrimerToken
		Función que obtiene el primer token de la cadena. Se utiliza la variable global cadena y posicion
		Returns:
			cadena[0] - Primer token de la cadena
*/
func PrimerToken() byte {
	posicion = 1
	return cadena[0]
}

/*
	 Function: SiguienteToken
		Función que obtiene el siguiente token de la cadena. Se utiliza la variable global cadena y posicion
		Returns:
			SiguienteToken() - Siguiente token de la cadena en caso de que exista un espacio en la expresión aritmética
			cadena[posicion-1] - Siguiente token de la cadena en caso de que no exista un espacio en la expresión aritmética
			; - Símbolo de fin de cadena

		See Also:
			<SiguienteToken>
*/
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

/*
	 Function: HacerMatch
		Función que verifica si el token de entrada es igual al token esperado. En caso de serlo, se obtiene el siguiente token de la cadena. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		Parameters:
			t - Token esperado
		See Also:
			<SiguienteToken>
*/
func HacerMatch(t byte) {
	if t == Token_Entrada {
		//Log = Log + ("Match: " + string(Token_Entrada) + " " + strconv.Itoa(posicion) + "\n")
		Token_Entrada = SiguienteToken()
	} else {
		Log = Log + ("error\n")
	}
}

/*
	 Function: SiguienteTokenID
		Función que obtiene el siguiente token de la cadena pero solo para los identificadores dado que no deberian tener espacios. Se utiliza la variable global cadena y posicion
		Returns:
			cadena[posicion-1] - Siguiente caracter del token de la expresión aritmética
			; - Símbolo de fin de cadena
*/
func SiguienteTokenID() byte {
	if posicion < len(cadena) {
		posicion = posicion + 1
		return cadena[posicion-1]
	} else {
		posicion = posicion + 1
		return ';'
	}
}

/*
	 Function: HacerMatchID
		Función que verifica si el token de entrada es igual al token esperado. En caso de serlo, se obtiene el siguiente token de la cadena. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		Parameters:
			t - Token esperado
		See Also:
			<SiguienteTokenID>
*/
func HacerMatchID(t byte) {
	if t == Token_Entrada {
		//Log = Log + ("Match: " + string(Token_Entrada) + " " + strconv.Itoa(posicion) + "\n")
		Token_Entrada = SiguienteTokenID()
	} else {
		Log = Log + ("error\n")
	}
}

//FUNCIONES DE LA GRAMATICA

/*
	 Function: expresion
		Función que analiza la gramática de la expresión aritmética. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<termino>
			<expresion_prima>
*/
func expresion() {
	termino()
	expresion_prima()
}

/*
	 Function: expresion_prima
		Función que analiza la gramática de la expresión aritmética. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<HacerMatch>
			<termino>
			<expresion_prima>
*/
func expresion_prima() {

	if Token_Entrada == '+' {
		preFija = preFija + "+"
		HacerMatch('+')

		termino()
		expresion_prima()
		postFija = postFija + "+"
	} else if Token_Entrada == '-' {
		preFija = preFija + "-"
		HacerMatch('-')

		termino()
		expresion_prima()
		postFija = postFija + "-"
	} else {
		//epsilon
	}
}

/*
	 Function: termino
		Función que analiza la gramática de la expresión aritmética en la regla de producción de termino. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<factor>
			<termino_prima>
*/
func termino() {
	factor()
	termino_prima()
}

/*
	 Function: termino_prima
		Función que analiza la gramática de la expresión aritmética en la regla de producción de termino_prima. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<HacerMatch>
			<factor>
			<termino_prima>
*/
func termino_prima() {
	if Token_Entrada == '*' {
		preFija = preFija + "*"
		HacerMatch('*')

		factor()

		termino_prima()
		postFija = postFija + "*"
	} else if Token_Entrada == '/' {
		preFija = preFija + "/"
		HacerMatch('/')

		factor()

		termino_prima()
		postFija = postFija + "/"
	} else if Token_Entrada == '%' {
		preFija = preFija + "%"
		HacerMatch('%')

		factor()

		termino_prima()
		postFija = postFija + "%"
	} else if Token_Entrada == '^' {
		preFija = preFija + "^"
		HacerMatch('^')

		factor()

		termino_prima()
		postFija = postFija + "^"
	} else {
		//epsilon
	}
}

/*
	 Function: factor
		Función que analiza la gramática de la expresión aritmética en la regla de producción de factor. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<HacerMatch>
			<expresion>
			<SiguienteToken>
			<expresion_prima>
			<termino_prima>
			<is_cov>
			<is_letter>
			<cov>
			<numero>
*/
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
			preFija = preFija + string(Token_Entrada)
			postFija = postFija + string(Token_Entrada)
			HacerMatch('@')
			numero()
		}
	} else {
		Log = Log + ("Error: Se esperaba un termino en la posición:" + strconv.Itoa(posicion) + "\n")
	}
}

/*
	 Function: cov
		Función que analiza la gramática de la expresión aritmética en la regla de producción de cov. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<HacerMatch>
			<identificador>
*/
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

/*
	 Function: identificador
		Función que analiza la gramática de la expresión aritmética en la regla de producción de identificador. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<letra>
			<identificador_prima>
*/
func identificador() {
	letra()
	identificador_prima()
}

/*
	 Function: identificador_prima
		Función que analiza la gramática de la expresión aritmética en la regla de producción de identificador_prima. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<is_letter>
			<letra>
			<identificador_prima>
			<SiguienteToken>
*/
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

/*
	 Function: letra
		Función que analiza la gramática de la expresión aritmética en la regla de producción de letra. Se utilizan las variables globales cadena, posicion, log y Token_Entrada
		See Also:
			<is_letter>
			<HacerMatchID>
*/
func letra() {
	if is_letter(Token_Entrada) {
		HacerMatchID(Token_Entrada)
	} else {
		Log = Log + ("Error: Se esperaba una letra en la posición:" + strconv.Itoa(posicion) + "\n")
	}
}

/*
	 Function: numero
		Función que analiza la gramática de la expresión aritmética en la regla de producción de numero.
		See Also:
			<digito>
			<numero_prima>
*/
func numero() {
	digito()
	numero_prima()
}

/*
	 Function: numero_prima
		Función que analiza la gramática de la expresión aritmética en la regla de producción de numero_prima.
		See Also:
			<is_digit>
			<digito>
			<numero_prima>
*/
func numero_prima() {
	if is_digit(Token_Entrada) {
		digito()
		numero_prima()
	} else {
		//epsilon
	}
}

/*
	 Function: digito
		Función que analiza la gramática de la expresión aritmética en la regla de producción de digito. Se utilizan las variables globales posicion, log y Token_Entrada
		See Also:
			<is_digit>
			<HacerMatch>
*/
func digito() {
	if is_digit(Token_Entrada) {
		preFija = preFija + string(Token_Entrada)
		postFija += string(Token_Entrada)
		HacerMatch(Token_Entrada)
	} else {
		Log = Log + ("Error: Se esperaba un digito, variable o constante en la posición:" + strconv.Itoa(posicion) + "\n")
	}
}

/*
	 Function: is_digit
		Función que verifica si el caracter es un digito
		Parameters:
			tokenEntrada - Caracter a verificar
		Returns:
			true - si el caracter es un digito
			false - en caso contrario
*/
func is_digit(tokenEntrada byte) bool {
	if tokenEntrada == '0' || tokenEntrada == '1' || tokenEntrada == '2' || tokenEntrada == '3' || tokenEntrada == '4' || tokenEntrada == '5' || tokenEntrada == '6' || tokenEntrada == '7' || tokenEntrada == '8' || tokenEntrada == '9' {
		return true
	} else {
		return false
	}
}

/*
	 Function: is_cov
		Función que verifica si el caracter es una constante o variable
		Parameters:
			tokenEntrada - Caracter a verificar
		Returns:
			true - si el caracter es una constante o variable
			false - en caso contrario
*/
func is_cov(tokenEntrada byte) bool {
	if tokenEntrada == '@' || tokenEntrada == '$' {
		return true
	} else {
		return false
	}
}

/*
	 Function: is_letter
		Función que verifica si el caracter es una letra
		Parameters:
			tokenEntrada - Caracter a verificar
		Returns:
			true - si el caracter es una letra
			false - en caso contrario
*/
func is_letter(tokenEntrada byte) bool {
	if unicode.IsLetter(rune(tokenEntrada)) {
		return true
	} else {
		return false
	}
}
