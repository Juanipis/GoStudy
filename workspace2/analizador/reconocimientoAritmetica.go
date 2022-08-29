package analizador

import (
	"fmt"
)

func RAritmetico(tabla [][]string) {
	var linea string
	for i := 0; i < len(tabla); i++ {
		if tabla[i][3] == "Operador aritmetico" && tabla[i][1] != linea {
			linea = tabla[i][1]
			fmt.Println(filter(tabla, i, linea))
		}
	}
}

func filter(tabla [][]string, pos int, linea string) []string {
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
	return []string{lineaCodigo, linea, posI, posF}
}
