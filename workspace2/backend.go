package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	analizador "workspace2/analizador"

	"github.com/gorilla/mux"
)

type Tesla struct {
	ProductName  string `json:"productName"`
	ProductPrice string `json:"productPrice"`
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	result := getTable()
	log.Println(result[0])
	//prueba := result[0]
	b, err := json.Marshal(result)

	/*
		json := simplejson.New()
		primes := []string{"hola", "mundo"}
		json.SetPath(primes)

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}*/
	w.Write(b)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
	log.Println(err)

}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", r))
}

func getTable() []analizador.Token {
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

	//var tablajson string
	var tokenTable1 []analizador.Token
	//tablajson = "{\n\t\"tabla1\": [\n"
	for _, elementTablaInt := range tablaIntermedia {
		for _, elementFila := range elementTablaInt {
			tokenTable1 = append(tokenTable1, elementFila)
			/*if len(elementFila.TypeToken) == 3 {
				elemento = "{\n\t\t\"nombre\":\"" + elementFila.TokenName + "\",\n\t\t\"linea\":\"" + strconv.Itoa(fila) + "\",\n\t\t\"columna\":\"" + strconv.Itoa(columna) + "\",\n\t\t\"tipo1\":\"" + elementFila.TypeToken[0] + "\",\n\t\t\"tipo2\":\"" + elementFila.TypeToken[1] + "\",\n\t\t\"tipo3\":\"" + elementFila.TypeToken[2] + "\"\t\t},\n"
			} else if len(elementFila.TypeToken) == 2 {
				elemento = "{\n\t\t\"nombre\":\"" + elementFila.TokenName + "\",\n\t\t\"linea\":\"" + strconv.Itoa(fila) + "\",\n\t\t\"columna\":\"" + strconv.Itoa(columna) + "\",\n\t\t\"tipo1\":\"" + elementFila.TypeToken[0] + "\",\n\t\t\"tipo2\":\"" + elementFila.TypeToken[1] + "\",\n\t\t\"tipo3\":\"\"" + "\n\t\t},\n"
			} else if len(elementFila.TypeToken) == 1 {
				elemento = "{\n\t\t\"nombre\":\"" + elementFila.TokenName + "\",\n\t\t\"linea\":\"" + strconv.Itoa(fila) + "\",\n\t\t\"columna\":\"" + strconv.Itoa(columna) + "\",\n\t\t\"tipo1\":\"" + elementFila.TypeToken[0] + "\",\n\t\t\"tipo2\":\"\"" + ",\n\t\t\"tipo3\":\"\"" + "\n\t\t},"
			}*/
		}
		//tablajson = tablajson + elemento
	}
	/*tablajson = tablajson[:len(tablajson)-1]
	tablajson = tablajson + "]\n}"*/
	return tokenTable1
}
