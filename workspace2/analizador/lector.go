package analizador

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaIntermedia [][]Token, tablaCorrespondencia map[string]string, tablaTokensGenerada []TablaTokens, tablatokens map[string]string) []TablaTokens {
	var stringAcumulado string
	//La linea a evaluar es un array de Tokens, los Tokens tienen un nombre y un array de tipos
	lineaEvaluar := []Token{}

	for i := 0; i < len(linea); i++ {
		caracterActual := string(linea[i])

		//Se verifica si el caracter actual es un separador en la tabla de simbolos
		tipoSeparador, isSeparator := tablaSimbolos[caracterActual]
		if isSeparator {
			//Algunos simbolos son de dos caracteres, revisemos si el siguiente caracter es un separador en conjunto con el actual
			if i+1 < len(linea) {
				caracterSiguiente := string(linea[i+1])
				//Se verifica si el caracter actual con el siguiente es un separador en la tabla de simbolos
				_, isSeparator := tablaSimbolos[caracterActual+caracterSiguiente]
				if isSeparator {
					caracterActual = caracterActual + caracterSiguiente
					i = i + 1
				}
			}
			//Revisa si el stringAcumulado acumulado es un simbolo de la tabla
			if stringAcumulado != "" {
				typeTokenTemp, isReserved := tablaSimbolos[stringAcumulado]
				if isReserved {
					//Si es un simbolo reservado, lo añade a la tabla de Simbolos con su respectivos tipos
					lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, ConversionTypeSimbolo(typeTokenTemp, tablaCorrespondencia)})
					tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{ObtencionToken(stringAcumulado, tablatokens), len(tablaTokensGenerada), stringAcumulado})
				} else {
					//Si el primer caracter del stringAcumulado es un arroba, significa que es una constante.
					if strings.HasPrefix(stringAcumulado, "@") {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Constante"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"SimboloConstante", len(tablaTokensGenerada), stringAcumulado})
						//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
					} else if strings.HasPrefix(stringAcumulado, "$") {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Variable"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"SimboloVariable", len(tablaTokensGenerada), stringAcumulado})
						//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
					} else {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Identificador", len(tablaTokensGenerada), stringAcumulado})
					}
				}
			}
			stringAcumulado = ""
			//Añade el separador a la tabla de Tokens con su respectivo tipo
			lineaEvaluar = append(lineaEvaluar, Token{caracterActual, ConversionTypeSimbolo(tipoSeparador, tablaCorrespondencia)})
			//tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{ObtencionToken(stringAcumulado, tablatokens), len(tablaTokensGenerada), stringAcumulado})
			test := TablaTokens{ObtencionToken(caracterActual, tablatokens), len(tablaTokensGenerada), caracterActual}
			tablaTokensGenerada = append(tablaTokensGenerada, test)
		} else {
			//No encontro un separador, lo concatena al stringAcumulado y sigue evaluando la linea actual
			stringAcumulado += caracterActual
		}
	}

	//Aca ya acabo la linea, se verifica si el stringAcumulado no esta vacio
	if stringAcumulado != "" {
		typeTokenTemp, isReserved := tablaSimbolos[stringAcumulado]
		if isReserved {
			lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, ConversionTypeSimbolo(typeTokenTemp, tablaCorrespondencia)})
			tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{ObtencionToken(stringAcumulado, tablatokens), len(tablaTokensGenerada), stringAcumulado})
		} else {
			if strings.HasPrefix(stringAcumulado, "@") {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Constante"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"SimboloConstante", len(tablaTokensGenerada), stringAcumulado})
			} else if strings.HasPrefix(stringAcumulado, "$") {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Variable"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"SimboloVariable", len(tablaTokensGenerada), stringAcumulado})
			} else {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Identificador", len(tablaTokensGenerada), stringAcumulado})
			}
		}
	}
	//Al terminar de recorrer la linea añade el conjunto de Tokens a la tabla final
	tablaIntermedia[numlinea] = lineaEvaluar
	return tablaTokensGenerada
}

func ConversionTypeSimbolo(tiposSimbolo []string, tablaCorrespondencia map[string]string) []string {
	valorReal := []string{}
	for i := 0; i < len(tiposSimbolo); i++ {
		tiposSimbolo, isPresent := tablaCorrespondencia[tiposSimbolo[i]]
		if isPresent {
			valorReal = append(valorReal, tiposSimbolo)
		}
	}
	return valorReal
}
func ObtencionToken(lexema string, tablatokens map[string]string) string {
	token, isPresent := tablatokens[lexema]
	if isPresent {
		return token
	} else {
		return ""
	}
}

func ContadorLineas(nombreArchivo string) int {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println(err)
	}
	defer archivo.Close()
	fileScanner := bufio.NewScanner(archivo)
	contadorLineas := 0
	for fileScanner.Scan() {
		contadorLineas++
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}
	return contadorLineas
}
