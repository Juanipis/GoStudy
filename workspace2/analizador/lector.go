package analizador

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaIntermedia [][]Token, tablaCorrespondencia map[string]string, tablaTokensGenerada []TablaTokens, tablatokens map[string][]string) []TablaTokens {
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
					aux := ObtencionToken(stringAcumulado, tablatokens)
					tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{aux[1], aux[0], stringAcumulado})
				} else {
					//Si el primer caracter del stringAcumulado es un arroba, significa que es una constante.
					if strings.HasPrefix(stringAcumulado, "@") {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Constante"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorConstante", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
						//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
					} else if strings.HasPrefix(stringAcumulado, "$") {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Variable"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorVariable", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
						//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
					} else {
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Identificador", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
					}
				}
			}
			stringAcumulado = ""
			if caracterActual == "<!" {
				stringAcumulado += "<"
				for i < len(linea) {
					if string(linea[i]) == "!" && i < len(linea)-1 && string(linea[i+1]) == ">" {
						stringAcumulado += "!>"
						i += 2
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Comentario"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Comentario", generarId(tablaTokensGenerada, len(tablatokens)), stringAcumulado})
						break
					} else {
						stringAcumulado += string(linea[i])
						i++
					}
				}
				stringAcumulado = ""
			} else {
				lineaEvaluar = append(lineaEvaluar, Token{caracterActual, ConversionTypeSimbolo(tipoSeparador, tablaCorrespondencia)})
				aux := ObtencionToken(caracterActual, tablatokens)
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{aux[1], aux[0], caracterActual})
			}
			//Añade el separador a la tabla de Tokens con su respectivo tipo
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
			aux := ObtencionToken(stringAcumulado, tablatokens)
			tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{aux[1], aux[0], stringAcumulado})
		} else {
			if strings.HasPrefix(stringAcumulado, "@") {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Constante"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorConstante", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
			} else if strings.HasPrefix(stringAcumulado, "$") {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Variable"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorVariable", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
			} else {
				lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador"}})
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Identificador", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
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
func ObtencionToken(lexema string, tablatokens map[string][]string) []string {
	token, isPresent := tablatokens[lexema]
	if isPresent {
		return token
	} else {
		return []string{"Error", "err"}
	}
}

func ExistenciaToken(lexema string, tablaTokens []TablaTokens, i int) string {
	existe := false
	var id string
	for Token := range tablaTokens {
		if lexema == tablaTokens[Token].LexemaGenerador {
			existe = true
			id = tablaTokens[Token].IdToken
			break
		}
	}
	if existe {
		return id
	} else {
		return generarId(tablaTokens, i)
	}
}

func generarId(tablaTokens []TablaTokens, i int) string {
	return strconv.Itoa(len(tablaTokens) + i)
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

func leerComentarios() {

}
