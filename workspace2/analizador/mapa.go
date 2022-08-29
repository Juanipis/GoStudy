package analizador

import (
	"bufio"
	"log"
	"os"
	"strings"
)

/*Generacion de estructura para manejo de simbolo
Se compone de un nombre y un arreglo de tipos, todos ellos como string
Su existencia es necesaria para la funcion principal de lectura*/
type Token struct {
	TokenName string
	TypeToken []string
}

/*Generacion de estructura para manejo de los tokens del codigo fuente
Se compone de un token definido por los desarrolladores, un idToken de este
y el lexema generador correspondiente (simbolo)
*/
type TablaTokens struct {
	Token           string
	IdToken         string
	LexemaGenerador string
}

/*Generacion de estructura para manejo de la tabla de simbolos que surge del analisis.
Se compone de aquellos datos asociados a las columnas (Tabla1) que van a ser visibles en el programa
*/
type FinalSimbol struct {
	Nombre      string `json:"name"`
	Linea       string `json:"line"`
	NumSimbFila string `json:"numSimbFila"`
	T1          string `json:"t1"`
	T2          string `json:"t2"`
	T3          string `json:"t3"`
}

//Metodo que crea una escructura de tipo map para el csv de TablaSimbolos
func CrearMapaSimbolos(path string, tablaSimbolos map[string][]string) {
	//Apertura del archivo y busqueda de errores
	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Error abriendo archivo : %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	//Se escanea cada uno de los elementos del csv y se separan por comas para el map.
	for fileScanner.Scan() {
		lista := strings.Split(fileScanner.Text(), ",")
		/*Llenado de mapa segun la cantidad de tipos que puede tener cada simbolo segun el csv.
		Llave: simbolo, valor: vector con la cantida de tipos que puede tener
		*/
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
		/*Llenado de mapa: Llave es el numero asignado en los tipos y el valor es su correspondiente
		Su utilidad se fundamenta en la traduccion de los numeros que tiene el mapa anterior (tablaSimbolos)
		para los tipos.
		*/
		tablaCorrespondencia[lista[0]] = lista[1]
	}
}

//Metodo que crea una estructura de tipo map para el csv de TablaTokens
func CrearMapaTokens(path string, tablaTokens map[string][]string) {
	//Apertura del archivo y busqueda de errores
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error abriendo archivo: %s", err)
	}
	fileScanner := bufio.NewScanner(file)

	//Se escanea cada uno de los elementos del csv y se separan por comas para el map.
	for fileScanner.Scan() {
		//Llenado de mapa: Llave es el token correspondiente y sus valores es un vector ligado a su id y el lexema que lo genera
		lista := strings.Split(fileScanner.Text(), ",")
		tablaTokens[lista[0]] = []string{lista[1], lista[2]}
	}
}
