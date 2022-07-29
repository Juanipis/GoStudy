package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type token struct {
	tokenName string
	typeToken []string
}

func crearMapa(path string, tablaSimbolos map[string][]string) {
	//Abre el archivo de entrada
	fd, error := os.Open(path)
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

//Luego limpia el string temporaeparador o un salto de linea que acabaria la instancia
func leer(linea string, numlinea int, tablaSimbolos map[string][]string, tablaFinal [][]token) {
	var stringAuxToken string
	lineaEvaluar := []token{}

	for i := 0; i < len(linea); i++ {
		caracterActual := string(linea[i])
		tipoSeparador, isSeparator := tablaSimbolos[caracterActual]
		//++
		//Algunos simbolos son de dos caracteres, revisemos si el siguiente caracter es un separador en conjunto con el actual
		if isSeparator {
			//Condiciones de separador doble
			if i+1 < len(linea) {
				caracterSiguiente := string(linea[i+1])
				_, isSeparator := tablaSimbolos[caracterActual+caracterSiguiente]
				if isSeparator {
					caracterActual = caracterActual + caracterSiguiente
					i = i + 1
				}
			}
			//Revisa si el stringAuxToken acumulado es un simbolo de la tabla
			typeTokenTemp, isReserved := tablaSimbolos[stringAuxToken]
			if isReserved && stringAuxToken != "" {
				//Si es un simbolo reservado, lo añade a la tabla de tokens con su respectivo tipo
				lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, typeTokenTemp})
			} else if stringAuxToken != "" {
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
	typeTokenTemp, isReserved := tablaSimbolos[stringAuxToken]
	if isReserved && stringAuxToken != "" {
		lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, typeTokenTemp})
	} else if stringAuxToken != "" {
		lineaEvaluar = append(lineaEvaluar, token{stringAuxToken, []string{"0"}})
	}
	//Al terminar de recorrer la linea añade el conjunto de tokens a la tabla de
	tablaFinal[numlinea] = lineaEvaluar
}

func contadorLineas(nombreArchivo string) int {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println(err)
	}
	defer archivo.Close()
	fileScanner := bufio.NewScanner(archivo)
	contadorLineas := 0

	//Se lee el archivo linea a linea
	for fileScanner.Scan() {
		contadorLineas++
	}
	//Se comprueba si hay un error
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}
	return contadorLineas
}

//Main
func main() {
	//0. Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	crearMapa("tblsimb.csv", tablaSimbolos)

	//1. Obtener cantidad de lineas
	nombreArchivo := "prog.messi"
	//fmt.Printf("Ingrese la ruta del archivo:")
	//fmt.Scan(&nombreArchivo)

	cantidadLineas := contadorLineas("prog.messi")

	tablaFinal := make([][]token, cantidadLineas)

	//
	//Ahora vamos a enviar las lineas a subrutinas de lectura go
	archivo, err := os.Open(nombreArchivo)
	//Se comprueba si hay un error
	if err != nil {
		fmt.Println(err)
	}
	defer archivo.Close()
	//Se crea un buffer de lectura
	fileScanner := bufio.NewScanner(archivo)
	i := 0
	for fileScanner.Scan() {
		leer(fileScanner.Text(), i, tablaSimbolos, tablaFinal)
		i++
	}

	for _, linea := range tablaFinal {
		fmt.Println(linea)
	}
}
