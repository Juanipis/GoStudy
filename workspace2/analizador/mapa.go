package analizador

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Class: Token
// Generacion de estructura para manejo de simbolo\nSu existencia es necesaria para la funcion principal de lectura
type Token struct {
	// Variable: TokenName
	// El nombre del token
	TokenName string
	// Variable: TypeToken
	// El arreglo de los tipos del token
	TypeToken []string
}

/*Generacion de estructura para manejo de los tokens del codigo fuente
Se compone de un token definido por los desarrolladores, un idToken de este
y el lexema generador correspondiente (simbolo)
*/

// Class: TablaTokens
// Generacion de estructura para manejo de los tokens del codigo fuente
type TablaTokens struct {
	// Variable: Token
	// El nombre del Token, este es definido por el desarrollador
	Token string
	// Variable: IdToken
	// El id del token, este es único segun el token
	IdToken string
	// Variable: Lexema
	// El lexema generado correspondiente al token ()
	LexemaGenerador string
}

/*Generacion de estructura para manejo de la tabla de simbolos que surge del analisis.
Se compone de aquellos datos asociados a las columnas (Tabla1) que van a ser visibles en el programa
*/

// Class: FinalSimbol
// Generacion de estructura para manejo de la tabla de simbolos que surge del analisis.\nSe compone de aquellos datos asociados a las columnas (Tabla1) que van a ser visibles en el programa
type FinalSimbol struct {
	// Variable: Nombre
	// El nombre del simbolo
	Nombre string `json:"name"`
	// Variable: Linea
	// La linea en la que se encuentra el simbolo
	Linea string `json:"line"`
	// Variable: NumSimbFila
	// La columna en la que se encuentra el simbolo
	NumSimbFila string `json:"numSimbFila"`
	// Variable: T1
	// El tipo 1 del simbolo
	T1 string `json:"t1"`
	// Variable: T2
	// El tipo 2 del simbolo
	T2 string `json:"t2"`
	// Variable: T3
	// El tipo 3 del simbolo
	T3 string `json:"t3"`
}

/*
Function: CrearMapaSimbolos

	Metodo que crea una escructura de tipo map para el csv de TablaSimbolos

	Parameters:

	   path - La ruta del archivo csv donde se encuentran los datos de la tabla de simbolos
	   tablaSimbolos - El mapa que se va a llenar con los datos del csv
*/
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

/*
Function: CrearMapaCorrespondencia

	Metodo que crea una escructura de tipo map para el csv de TablaCorrespondencia

	Parameters:

	   path - La ruta del archivo csv donde se encuentran los datos de la tabla de correspondencia
	   tablaCorrespondencia - El mapa que se va a llenar con los datos del csv
*/
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

/*
Function: CrearMapaTokens

	Metodo que crea una escructura de tipo map para el csv de TablaTokens

	Parameters:

	   path - La ruta del archivo csv donde se encuentran los datos de la tabla de tokens
	   tablaTokens - El mapa que se va a llenar con los datos del csv
*/
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
