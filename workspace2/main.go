package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	analizador "workspace2/analizador"
)

func main() {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	analizador.CrearMapa2("TablaSimbolos.csv", tablaSimbolos)
	tablaCorrespondencia := make(map[string]string)
	analizador.CrearMapaCorrespondencia("TablaCorrespondencia.csv", tablaCorrespondencia)
	Mapatokens := make(map[string]string)
	analizador.CrearMapaTokens("TablaTokens.csv", Mapatokens)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"

	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := analizador.ContadorLineas(nombreArchivo)

	//Se crea la tabla intermedia que alimenta a la final de tokens
	tablaIntermedia := make([][]analizador.Token, cantidadLineas)

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println(err)
	}

	defer archivo.Close()
	fileScanner := bufio.NewScanner(archivo)
	//Se inicializa la variable que envia la fila que se esta analizando
	i := 0

	//Se lee el archivo linea a linea enviando al metodo leer

	for fileScanner.Scan() {
		tabla2 = analizador.Leer(fileScanner.Text(), i, tablaSimbolos, tablaIntermedia, tablaCorrespondencia, tabla2, Mapatokens)
		i++
	}

	//Se creará la tabla que se exportara como csv
	final := [][]string{}
	final = append(final, []string{"NOMBRE", "LINEA", "# SIMBOLO EN FILA", "TIPO 1", "TIPO 2", "TIPO 3"})

	for f, linea := range tablaIntermedia {
		for c, token := range linea {
			var tipos [4]string
			copy(tipos[:], token.TypeToken)
			//Se agrega simbolo en la linea f y posicion c a la tabla final con su respectiva posicion y tipos
			final = append(final, []string{token.TokenName, strconv.Itoa(f), strconv.Itoa(c), tipos[0], tipos[1], tipos[2]})
		}
	}

	analizador.RAritmetico(final)
	// for _, elemento := range tabla2 {
	// 	fmt.Println(elemento)
	// }

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
