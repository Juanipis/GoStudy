package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	analizador "workspace2/analizador"
)

func main() {
	tablaSimbolos := make(map[string][]string)
	analizador.CrearMapa2("../TablaSimbolos.csv", tablaSimbolos)

	nombreArchivo := "../prog.messi"

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
	//Para usar subrutinas creamos un sync
	var wg sync.WaitGroup

	for fileScanner.Scan() {
		wg.Add(1)
		go analizador.Leer(fileScanner.Text(), i, tablaSimbolos, tablaIntermedia, &wg)
		i++
	}

	wg.Wait()

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

	//Ahora vamos a guardar la tabla finañ
	f, _ := os.Create("../AnalizadorLexicoGrafico.csv")
	w := bufio.NewWriter(f)

	for _, tokensito := range final {
		var bufer string
		for _, index := range tokensito {
			bufer += index + ","
		}
		w.WriteString(bufer + "\n")
	}
	w.Flush()

}
