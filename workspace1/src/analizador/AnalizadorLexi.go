package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type token struct {
	tokenName string
	typeToken []string
}

//Función que genera mapa para reconocimiento de simbolos
func crearMapa(path string, tablaSimbolos map[string][]string) {
	//Abre el archivo de entrada
	fd, error := os.Open(path)
	//Se comprueba si hay un error
	if error != nil {
		fmt.Println(error)
	}
	defer fd.Close()

	// Lee el archivo
	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()
	//Se comprueba si hay un error en la lectura
	if error != nil {
		fmt.Println(error)
	}
	// Crea el mapa de simbolos, el cual tiene como clave el simbolo y como valor un array de strings con los tipos de simbolos
	for _, lista := range records {
		if lista[3] != "" {
			tablaSimbolos[lista[0]] = []string{lista[1], lista[2], lista[3]}
		} else if lista[2] != "" {
			tablaSimbolos[lista[0]] = []string{lista[1], lista[2]}
		} else {
			tablaSimbolos[lista[0]] = []string{lista[1]}
		}
	}
}

//Función principal de clasificacion de tokens
//Se le pasa la linea del archivo a analizar, el numero de esta linea, el mapa de simbolos y la tabla final
func leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaFinal [][]token, wg *sync.WaitGroup) {
	//Se crea un string temporal para almacenar los tokens
	var stringAuxToken string
	//La linea a evaluar es un array de tokens, los tokens tienen un nombre y un array de tipos
	lineaEvaluar := []token{}

	//Ciclo principal para analisis de la linea, recorre desde 0 hasta la longitud de la linea
	for i := 0; i < len(linea); i++ {
		//Se obtiene el caracter actual de la linea a evaluar
		caracterActual := string(linea[i])

		//Se verifica si el caracter actual es un separador en la tabla de simbolos
		tipoSeparador, isSeparator := tablaSimbolos[caracterActual]
		if isSeparator {
			//Algunos simbolos son de dos caracteres, revisemos si el siguiente caracter es un separador en conjunto con el actual
			if i+1 < len(linea) {
				//Asignacion de caracter siguiente si cumple condicion
				caracterSiguiente := string(linea[i+1])
				//Se verifica si el caracter actual con el siguiente es un separador en la tabla de simbolos
				_, isSeparator := tablaSimbolos[caracterActual+caracterSiguiente]
				//Si es en conjunto un separador, la variable caracter actual se combina
				if isSeparator {
					caracterActual = caracterActual + caracterSiguiente
					i = i + 1
				}
			}

			//Revisa si el stringAuxToken acumulado es un simbolo de la tabla
			if stringAuxToken != "" {
				typeTokenTemp, isReserved := tablaSimbolos[stringAuxToken]
				if isReserved {
					//Si es un simbolo reservado, lo añade a la tabla de tokens con su respectivo tipo
					lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, typeTokenTemp})
				} else {
					//Si el primer caracter del stringAcumulado es un arroba, significa que es una constante.
					if strings.HasPrefix(stringAuxToken, "@") {
						lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador", "Constante"}})
						//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
					} else if strings.HasPrefix(stringAuxToken, "$") {
						lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador", "Variable"}})
						//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
					} else {
						lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador"}})
					}
				}
			}
			//Limpia el stringAcumulado
			stringAuxToken = ""
			//Añade el separador a la tabla de tokens con su respectivo tipo
			lineaEvaluar = append(lineaEvaluar, token{caracterActual, tipoSeparador})
		} else {
			//No encontro un separador, lo concatena al stringAcumulado y sigue evaluando la linea actual
			stringAuxToken += caracterActual
		}
	}
	//Aca ya se acabo la linea, se verifica si el stringAcumulado no esta vacio
	if stringAuxToken != "" {
		typeTokenTemp, isReserved := tablaSimbolos[stringAuxToken]
		if isReserved {
			lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, typeTokenTemp})
		} else {
			if strings.HasPrefix(stringAuxToken, "@") {
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador", "Constante"}})
				//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
			} else if strings.HasPrefix(stringAuxToken, "$") {
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador", "Variable"}})
				//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
			} else {
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"Identificador"}})
			}
		}
	}
	//Al terminar de recorrer la linea añade el conjunto de tokens a la tabla de
	tablaFinal[numlinea] = lineaEvaluar

	wg.Done()
}

//Debido a que se van a utilizar subrutinas, se debe saber la cantidad de lineas que se van a analizar
func contadorLineas(nombreArchivo string) int {
	//Se abre el archivo
	archivo, err := os.Open(nombreArchivo)
	//Se verifica si hubo un error al abrir el archivo
	if err != nil {
		fmt.Println(err)
	}
	//Se cierra el archivo
	defer archivo.Close()
	//Se crea un buffer para leer el archivo
	fileScanner := bufio.NewScanner(archivo)
	//Se inicializa la variable que contara las lineas
	contadorLineas := 0
	//Se lee el archivo linea a linea
	for fileScanner.Scan() {
		//Se aumenta el contador de lineas
		contadorLineas++
	}
	//Se comprueba si hay un error
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}
	//Se retorna la cantidad de lineas
	return contadorLineas
}

//Main
func main() {
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	crearMapa("tblsimb.csv", tablaSimbolos)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"
	//fmt.Printf("Ingrese la ruta del archivo:")
	//fmt.Scan(&nombreArchivo)
	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := contadorLineas(nombreArchivo)
	//Se crea la tabla final de tokens
	tablaFinal := make([][]token, cantidadLineas)

	//Se abre el archivo fuente
	archivo, err := os.Open(nombreArchivo)
	//Se comprueba si hay un error
	if err != nil {
		fmt.Println(err)
	}
	//Se cierra el archivo al final
	defer archivo.Close()
	//Se crea un buffer de lectura
	fileScanner := bufio.NewScanner(archivo)
	//Se inicializa la variable que envia la fila que se esta analizando
	i := 0
	//Se lee el archivo linea a linea enviando al metodo leer

	//Para usar subrutinas creamos un sync
	var wg sync.WaitGroup

	for fileScanner.Scan() {
		wg.Add(1)
		go leer(fileScanner.Text(), i, tablaSimbolos, tablaFinal, &wg)
		i++
	}

	wg.Wait()
	//fmt.Println(tablaFinal)

	final := [][]string{}
	final = append(final, []string{"NOMBRE", "LINEA", "# SIMBOLO EN FILA", "TIPO 1", "TIPO 2", "TIPO 3"})

	for f, linea := range tablaFinal {
		for c, token := range linea {
			var tipos [3]string
			for i, tipo := range token.typeToken {
				tipos[i] = tipo
			}
			final = append(final, []string{token.tokenName, strconv.Itoa(f), strconv.Itoa(c), tipos[0], tipos[1], tipos[2]})
		}
	}

	csvFile, err := os.Create("AnalizadorLexicoGrafico.csv")
	if err != nil {
		fmt.Println(err)
	}

	csvwriter := csv.NewWriter(csvFile)
	for _, linea := range final {
		csvwriter.Write(linea)
	}
	csvwriter.Flush()
	csvFile.Close()
}
