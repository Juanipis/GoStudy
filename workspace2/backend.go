package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	analizador "workspace2/analizador"

	"github.com/gorilla/mux"
)

type finalSimbol struct {
	Nombre      string `json:"name"`
	Linea       string `json:"line"`
	NumSimbFila string `json:"numSimbFila"`
	T1          string `json:"t1"`
	T2          string `json:"t2"`
	T3          string `json:"t3"`
}

func Tabla1(w http.ResponseWriter, r *http.Request) {
	result := getTable1()
	b, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func Tabla2(w http.ResponseWriter, r *http.Request) {
	result := getTable2()
	b, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func Tabla3(w http.ResponseWriter, r *http.Request) {
	result := getTable3()
	b, _ := json.Marshal(result)
	fmt.Println(b)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	os.Remove("prog.messi")
	f, err := os.OpenFile("prog.messi", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
	_, _ = io.Copy(f, file)
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/tabla1", Tabla1).Methods("GET")
	r.HandleFunc("/tabla2", Tabla2).Methods("GET")
	r.HandleFunc("/tabla3", Tabla3).Methods("GET")
	r.HandleFunc("/file", UploadFile).Methods("POST")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", r))
}

func getTable1() []finalSimbol {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
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

	final2 := []finalSimbol{}
	for f, linea := range tablaIntermedia {
		for c, token := range linea {
			var tipos [4]string
			copy(tipos[:], token.TypeToken)
			//Se agrega simbolo en la linea f y posicion c a la tabla final con su respectiva posicion y tipos
			final2 = append(final2, finalSimbol{token.TokenName, strconv.Itoa(f), strconv.Itoa(c), tipos[0], tipos[1], tipos[2]})
		}

	}
	return final2
}

func getTable2() []analizador.TablaTokens {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
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
	return tabla2
}

func getTable3() []analizador.AritemticaStruct {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
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
	tablaAritmetica := analizador.RAritmetico(final)
	return tablaAritmetica
}
