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
	var tempToken string
	lineaTokens := []token{}

	for i := 0; i < len(linea); i++ {
		lineaActual := string(linea[i])
		typeTokenSeparator, isSeparator := tablaSimbolos[lineaActual]
		//if,hola//Encontro un separador
		if isSeparator {
			//Revisa si el tempToken acumulado es un simbolo reservado
			typeTokenTemp, isReserved := tablaSimbolos[tempToken]
			if isReserved {
				//Si es un simbolo reservado, lo añade a la tabla de tokens con su respectivo tipo
				lineaTokens = append(lineaTokens, token{tempToken, typeTokenTemp})
			} else {
				//si no es un simbolo reservado, lo añade a la tabla de tokens con tipo 0==Identificador
				lineaTokens = append(lineaTokens, token{tempToken, []string{"0"}})
			}
			//Limpia el tempToken
			tempToken = ""
			//Añade el separador a la tabla de tokens con su respectivo tipo
			lineaTokens = append(lineaTokens, token{lineaActual, typeTokenSeparator})
		} else {
			//No encontro un separador, lo concatena al tempToken y sigue leyendo
			tempToken += lineaActual
		}
	}
	//Al terminar de recorrer la linea añade el conjunto de tokens a la tabla de
	tablaFinal[fila] = lineaTokens
}
