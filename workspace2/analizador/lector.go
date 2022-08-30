package analizador

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Metodo principal para el analizador lexicografico
Funcionalidad: lectura del codigo fuente que se envia por ruta.
Valor de retorno: La tabla de tokens luego del analisis.
Parametros: linea: corresponde a aquella del archivo a ser analizada.
numlinea: identifica el numero de la linea con la que se esta trabajando.
tablaSimbolos: mapa que ayuda a saber si tenemos un simbolo reservado o no.
tablaIntermedia: aquella que se genera con todos los simbolos evaluados del archivo.
tablaCorrespondencia: mapa que ayuda a identificar el significado de los numeros que se encuentran en tipo del simbolo.
tablaTokensGenerada: tabla que se devuelve en el metodo con token, idToken y lexema.
tablatokens: mapa que ayuda a conseguir el token a partir del lexema (simbolo).
*/
func Leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaIntermedia [][]Token, tablaCorrespondencia map[string]string, tablaTokensGenerada []TablaTokens, tablatokens map[string][]string) []TablaTokens {
	var stringAcumulado string
	//La linea a evaluar es un array de Simbolos, los Simbolos tienen un nombre y un array de tipos
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
			//Revisa si el stringAcumulado que se tiene es un simbolo de la tabla
			if stringAcumulado != "" {
				//Se verifica si el stringAcumulado es una palabra en el mapa de simbolos
				typeTokenTemp, isReserved := tablaSimbolos[stringAcumulado]
				if isReserved {
					//Si es un simbolo reservado, lo añade a la tabla de Simbolos con su respectivos tipos
					lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, ConversionTypeSimbolo(typeTokenTemp, tablaCorrespondencia)})
					//Auxiliar que almacena para un lexema existente en la tabla de token, su token y idToken correspondiente.
					aux := ObtencionToken(stringAcumulado, tablatokens)
					//Se mete en la tabla final de tokens la informacion asociada a este
					tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{aux[1], aux[0], stringAcumulado})
				} else {
					//Si el primer caracter del stringAcumulado es un arroba, significa que es una constante.
					if strings.HasPrefix(stringAcumulado, "@") {
						//Se ingresa el identificador constante a tabla de simbolos y de tokens.
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Constante"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorConstante", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
						//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
					} else if strings.HasPrefix(stringAcumulado, "$") {
						//Se ingresa el identificador variable a tabla de simbolos y de tokens.
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador", "Variable"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"IdentificadorVariable", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
						//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
					} else {
						//Se ingresa el identificador a tabla de simbolos y de tokens.
						lineaEvaluar = append(lineaEvaluar, Token{stringAcumulado, []string{"Identificador"}})
						tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{"Identificador", ExistenciaToken(stringAcumulado, tablaTokensGenerada, len(tablatokens)), stringAcumulado})
					}
				}
			}

			//Se vacia el stringAcumulado dado a que ya se inserto en las tablas correspondientes.
			stringAcumulado = ""
			//Verificacion del caso de los comentarios para que se identifique toda la linea asociada a este.
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
				//Se vacia nuevamente el stringAcumulado
				stringAcumulado = ""
				//Si no es un comentario se procede a ingresarlo como un separador de cualquiera que se tiene.
			} else {
				lineaEvaluar = append(lineaEvaluar, Token{caracterActual, ConversionTypeSimbolo(tipoSeparador, tablaCorrespondencia)})
				aux := ObtencionToken(caracterActual, tablatokens)
				tablaTokensGenerada = append(tablaTokensGenerada, TablaTokens{aux[1], aux[0], caracterActual})
			}

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
	//Al terminar de recorrer la linea añade el conjunto de Tokens a la tabla intermedia
	tablaIntermedia[numlinea] = lineaEvaluar
	return tablaTokensGenerada
}

/* Metodo que convierte los tipos numericos contenidos en un array de simbolos a
su correspondiente significado en el mapa de correspondencia, con el fin de ingresarlo a la tabla final.
Parametros: array de simbolos y de correspondencia
Valor retorno: un vector de strings con el significado de todos los numeros asociados al simbolo
*/
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

/* Metodo que obtiene el token y el idToken correspondiente a cualquier tipo excepto identificadores.
Parametros: lexema a encontrar y mapa de tokens.
Valor de retorno: array con el token y el idToken para cada lexema.
*/
func ObtencionToken(lexema string, tablatokens map[string][]string) []string {
	token, isPresent := tablatokens[lexema]
	if isPresent {
		return token
	} else {
		return []string{"Error", "err"}
	}
}

/* Metodo que comprueba si un identificador ya fue inicializado anteriormente para asignarle
su idToken asociado o si no se le genera uno nuevo.
Parametros: lexema que contiene el token y id, la tabla de tokens final, i como longitud de la tabla
Valor retorno: string con el id del identificador evaluado.
*/
func ExistenciaToken(lexema string, tablaTokens []TablaTokens, lenTbTknGen int) string {
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
		return generarId(tablaTokens, lenTbTknGen)
	}
}

// Metodo que genera un id unico para el identificador, se aumenta desde el lenTbTknGen
func generarId(tablaTokens []TablaTokens, lenTbTknGen int) string {
	return strconv.Itoa(len(tablaTokens) + lenTbTknGen)
}

// Metodo que cuenta el total de lineas del codigo fuente, uso en main.
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
