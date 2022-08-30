package analizador

/*Class: AritemticaStruct
Estructura que guarda toda la informacion asociada a la expresion aritmetica y que debera ser almacenada
*/
type AritemticaStruct struct {
	// Variable: Expresion
	// Expresion aritmetica
	Expresion string `json:"exp"`
	// Variable: Linea
	// Linea del codigo en la que se encuentra la expresion aritmetica
	Linea string `json:"linea"`
	// Variable: SimInicial
	// Ubicacion del simbolo inicial de la expresion aritmetica en la linea de codigo correspondiente
	SimInicial string `json:"simInicio"`
	// Variable: SimInicial
	// Ubicacion del simbolo final de la expresion aritmetica en la linea de codigo correspondiente
	SimFinal string `json:"simFinal"`
}

/*Function: RAritmetico
Metodo principal que analiza las expresiones aritmeticas existentes dentro de una linea

Parameters:
    tabla - Matriz que simboliza a la tabla final de simbolos.

Returns:
    Arreglo de AritmeticaStruct, que representa todas las expresiones aritmeticas encontradas en el codigo junto con su ubicacion.

*/

func RAritmetico(tabla [][]string) []AritemticaStruct {
	var linea string
	var tablaAritmetica []AritemticaStruct
	for i := 0; i < len(tabla); i++ {
		if tabla[i][3] == "Operador aritmetico" && tabla[i][1] != linea {
			linea = tabla[i][1]
			tablaAritmetica = append(tablaAritmetica, filter(tabla, i, linea))
		}
	}
	return tablaAritmetica
}

/*Function: filter
Metodo que filtra las expresiones aritmeticas una vez han sido detectadas

Parameters:
    tabla - Matriz que simboliza a la tabla final de simbolos.
	pos - Posicion de un operador aritmetico en la tabla.
	linea - Linea en la que se encuentra el operador y expresion aritmetica.

Returns:
    AritemticaStruct, que representa la expresion aritmetica encontrada en el codigo junto con su ubicacion.

*/

func filter(tabla [][]string, pos int, linea string) AritemticaStruct {
	for {
		in := false
		pos--
		for i := 3; i < 6; i++ {
			if tabla[pos][i] == "Variable" || tabla[pos][i] == "Constante" {
				in = true
			}
		}
		if !in && !(tabla[pos][0] == "(" || tabla[pos][0] == ")" || tabla[pos][0] == " ") {
			if tabla[pos+1][0] == " " {
				pos = pos + 2
			} else {
				pos++
			}
			break
		}
	}
	posI := tabla[pos][2]
	var lineaCodigo string
	for {
		in := false
		for i := 3; i < 6; i++ {
			if tabla[pos][i] == "Variable" || tabla[pos][i] == "Constante" || tabla[pos][i] == "Operador aritmetico" {
				in = true
			}
		}
		if !in && !(tabla[pos][0] == "(" || tabla[pos][0] == ")" || tabla[pos][0] == " ") {
			if string(lineaCodigo[len(lineaCodigo)-1]) == " " {
				lineaCodigo = lineaCodigo[:len(lineaCodigo)-1]
				pos = pos - 2
			} else {
				pos--
			}
			break
		}
		lineaCodigo = lineaCodigo + tabla[pos][0]
		pos++
	}
	posF := tabla[pos][2]

	return AritemticaStruct{lineaCodigo, linea, posI, posF}
}
