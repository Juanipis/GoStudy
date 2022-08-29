package analizador

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Token struct {
	TokenName string
	TypeToken []string
}

type TablaTokens struct {
	Token           string
	IdToken         int
	LexemaGenerador string
}

//Metodo que crea una escructura de tipo map para el csv de TablaSimbolos
func CrearMapa2(path string, tablaSimbolos map[string][]string) {
	//Apertura del archivo y busqueda de errores
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error abriendo archivo : %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	//Se escanea cada uno de los elementos del csv y se separan por comas para el map.
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), ",")
		if lista[3] != "" {
			tablaSimbolos[lista[0]] = []string{lista[1], lista[2], lista[3]}
		} else if lista[2] != "" {
			tablaSimbolos[lista[0]] = []string{lista[1], lista[2]}
		} else {
			tablaSimbolos[lista[0]] = []string{lista[1]}
		}
	}
}

//Metodo que crea una escructura de tipo map para el csv de TablaCorrespondencia
func CrearMapaCorrespondencia(path string, tablaCorrespondencia map[string]string) {
	//Apertura del archivo y busqueda de errores
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error abriendo archivo : %s", err)
	}
	fileScanner := bufio.NewScanner(file)

	//Se escanea cada uno de los elementos del csv y se separan por comas para el map.
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), ",")
		tablaCorrespondencia[lista[0]] = lista[1]
	}
}

//Metodo que crea una estructura de tipo map para el csv de TablaTokens
func CrearMapaTokens(path string, tablaTokens map[string]string) {
	//Apertura del archivo y busqueda de errores
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error abriendo archivo: %s", err)
	}
	fileScanner := bufio.NewScanner(file)

	//Se escanea cada uno de los elementos del csv y se separan por comas para el map.
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), ",")
		tablaTokens[lista[0]] = lista[1]
	}
}

////////////////////////////////////////////////////////////////
func CrearMapa(path string, tablaSimbolos map[string][]string) {
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
