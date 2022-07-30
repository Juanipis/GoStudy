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
	fd, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
	}
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

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
func leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaIntermedia [][]token, wg *sync.WaitGroup) {
	var stringAcumulado string
	//La linea a evaluar es un array de tokens, los tokens tienen un nombre y un array de tipos
	lineaEvaluar := []token{}

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
					//Si es un simbolo reservado, lo añade a la tabla de tokens con su respectivo tipo
					lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, typeTokenTemp})
				} else {
					//Si el primer caracter del stringAcumulado es un arroba, significa que es una constante.
					if strings.HasPrefix(stringAcumulado, "@") {
						lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador", "Constante"}})
						//Si el primer caracter del stringAcumulado es un simbolo de pesos, significa que es una variable.
					} else if strings.HasPrefix(stringAcumulado, "$") {
						lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador", "Variable"}})
						//Si no cumple ninguno de los dos anteriores, se clasifica como identificador unicamente.
					} else {
						lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador"}})
					}
				}
			}

			stringAcumulado = ""
			//Añade el separador a la tabla de tokens con su respectivo tipo
			lineaEvaluar = append(lineaEvaluar, token{caracterActual, tipoSeparador})
		} else {
			//No encontro un separador, lo concatena al stringAcumulado y sigue evaluando la linea actual
			stringAcumulado += caracterActual
		}
	}

	//Aca ya acabo la linea, se verifica si el stringAcumulado no esta vacio
	if stringAcumulado != "" {
		typeTokenTemp, isReserved := tablaSimbolos[stringAcumulado]
		if isReserved {
			lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, typeTokenTemp})
		} else {
			if strings.HasPrefix(stringAcumulado, "@") {
				lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador", "Constante"}})
			} else if strings.HasPrefix(stringAcumulado, "$") {
				lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador", "Variable"}})
			} else {
				lineaEvaluar = append(lineaEvaluar, token{stringAcumulado, []string{"Identificador"}})
			}
		}
	}
	//Al terminar de recorrer la linea añade el conjunto de tokens a la tabla final
	tablaIntermedia[numlinea] = lineaEvaluar

	wg.Done()
}

//Debido a que se van a utilizar subrutinas, se debe saber la cantidad de lineas que se van a analizar
func contadorLineas(nombreArchivo string) int {
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

//Funcion Main
func main() {
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	crearMapa("TablaSimbolos.csv", tablaSimbolos)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"

	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := contadorLineas(nombreArchivo)

	//Se crea la tabla intermedia que alimenta a la final de tokens
	tablaIntermedia := make([][]token, cantidadLineas)

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println(err)
	}

	defer archivo.Close()
	fileScanner := bufio.NewScanner(archivo)
	//Se inicializa la variable que envia la fila que se esta analizando
	i := 0

	//Se lee el archivo linea a linea enviando al metodo leer
	//Para usar subrutinas creamos un sync
	var wg sync.WaitGroup

	for fileScanner.Scan() {
		wg.Add(1)
		go leer(fileScanner.Text(), i, tablaSimbolos, tablaIntermedia, &wg)
		i++
	}

	wg.Wait()

	//Se creará la tabla que se exportara como csv
	final := [][]string{}
	final = append(final, []string{"NOMBRE", "LINEA", "# SIMBOLO EN FILA", "TIPO 1", "TIPO 2", "TIPO 3"})

	for f, linea := range tablaIntermedia {
		for c, token := range linea {
			var tipos [4]string
			for i, tipo := range token.typeToken {
				tipos[i] = tipo
			}
			//Se agrega simbolo en la linea f y posicion c a la tabla final con su respectiva posicion y tipos
			final = append(final, []string{token.tokenName, strconv.Itoa(f), strconv.Itoa(c), tipos[0], tipos[1], tipos[2]})
		}
	}

	//Se exporta la tabla final como csv
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
