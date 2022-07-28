package analizador

import (
	"encoding/csv"
	"fmt"
	"os"
)

type token struct {
	tokenName string
	typeToken []string
}

var tablaSimbolos map[string][]string
var tablaFinal [][]token
var palabra string

func crearMapa(path string) {
	//Abre el archivo de entrada
	fd, error := os.Open("path")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println("Successfully opened the CSV file")
	defer fd.Close()

	// Lee el archivo
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	if error != nil {
		fmt.Println(error)
	}

	// Crea el mapa de tokens
	for _, lista := range records {
		if lista[2] == "" {
			tablaSimbolos[lista[0]] = []string{lista[1]}
		} else {
			tablaSimbolos[lista[0]] = []string{lista[1], lista[2]}
		}
	}
}

//Obtiene una linea del archivo de entrada.
//Lee caracter a caracter buscando un separador
//Los separadores estan guardados en un arreglo de runes
//Va almacenando el token en un string temporal mientras no encuentre separadores
//Cuando encuentra un separador, lo envia a la función reconocimiento
//Luego limpia el string temporaeparador o un salto de linea que acabaria la instancia
func leer(linea string, fila int) {
	var stringAuxToken string
	lineaEvaluar := []token{}

	for i := 0; i < len(linea); i++ {
		caracterActual := string(linea[i])
		tipoSeparador, isSeparator := tablaSimbolos[caracterActual]
		//++
		//Algunos simbolos son de dos caracteres, revisemos si el siguiente caracter es un separador en conjunto con el actual
		if isSeparator {
			if i+1 < len(linea) {
				caracterSiguiente := string(linea[i+1])
				_, isSeparator := tablaSimbolos[caracterActual+caracterSiguiente]
				if isSeparator {
					caracterActual = caracterActual + caracterSiguiente
					i = i + 2
				}
			}
			//Revisa si el stringAuxToken acumulado es un simbolo de la tabla
			typeTokenTemp, isReserved := tablaSimbolos[stringAuxToken]
			if isReserved {
				//Si es un simbolo reservado, lo añade a la tabla de tokens con su respectivo tipo
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, typeTokenTemp})
			} else {
				//si no es un simbolo reservado, lo añade a la tabla de tokens con tipo 0==Identificador
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"0"}})
			}
			//Limpia el stringAuxToken
			stringAuxToken = ""
			//Añade el separador a la tabla de tokens con su respectivo tipo
			lineaEvaluar = append(lineaEvaluar, token{caracterActual, tipoSeparador})
		} else {
			//No encontro un separador, lo concatena al stringAuxToken y sigue leyendo
			stringAuxToken += caracterActual
		}
	}
	//Al terminar de recorrer la linea añade el conjunto de tokens a la tabla de
	tablaFinal[fila] = lineaEvaluar
}
