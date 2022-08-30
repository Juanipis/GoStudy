package analizador

/* Estructura que guarda toda la informacion asociada a la expresion aritmetica y que debera ser almacenada
 */
type AritemticaStruct struct {
	Expresion  string `json:"exp"`
	Linea      string `json:"linea"`
	SimInicial string `json:"simInicio"`
	SimFinal   string `json:"simFinal"`
}

/*Metodo principal que analiza las expresiones aritmeticas existentes dentro de una linea
Parametros: Matriz que simboliza a la tabla final de simbolos
Valor de retorno: un array de la estructura aritmetica
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
