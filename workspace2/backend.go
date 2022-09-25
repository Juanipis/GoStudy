package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	analizador "workspace2/analizador"
	automatas "workspace2/automatas"

	"github.com/gorilla/mux"
)

/*
Function: Tabla1

	Metodo que recibe la peticion GET /tabla1 y devuelve la tabla de simbolos

	Parameters:

	   w - Escribe los datos a devolver
	   r - Solicitud entrante.
*/
func Tabla1(w http.ResponseWriter, r *http.Request) {
	result := getTable1()
	b, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

/*
Function: Tabla2

	Metodo que recibe la peticion GET /tabla2 y devuelve la tabla de tokens

	Parameters:

	   w - Escribe los datos a devolver
	   r - Solicitud entrante.
*/
func Tabla2(w http.ResponseWriter, r *http.Request) {
	result := getTable2()
	b, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

/*
Function: Tabla3

	Metodo que recibe la peticion GET /tabla3 y devuelve la tabla aritmetica

	Parameters:

	   w - Escribe los datos a devolver
	   r - Solicitud entrante.
*/
func Tabla3(w http.ResponseWriter, r *http.Request) {
	result := getTable3()
	b, _ := json.Marshal(result)
	fmt.Println(b)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

/* Function: UploadFile

   Metodo que recibe los archivos entrantes a ser analizados

   Parameters:

      w - escribe los datos recibidos en un archivo prog.messi
      r - Solicitud entrante.
*/

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

type ExprAritmetica struct {
	Exp string `json:"exp"`
}
type ResultAritmetica struct {
	Result bool   `json:"result"`
	Log    string `json:"log"`
}

func AnalizarExpresionAritmetica(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newExp ExprAritmetica
	json.Unmarshal(reqBody, &newExp)

	newData, err := json.Marshal(newExp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
		//Cuando tengamos listo
		result, log := automatas.Run(string(newData))
		arit := ResultAritmetica{Result: result, Log: log}
		//arit := ResultAritmetica{Result: true, Log: "Mathc: 1\n Match2"}
		json.NewEncoder(w).Encode(arit)
	}
}

/*
Function: main

	Metodo principal que abre el servidor y se mantine activo a la escucha de peticiones
*/
func main() {
	r := mux.NewRouter()
	// Las rutas consisten de un path y una funcion para enviar o recibir datos
	r.HandleFunc("/tabla1", Tabla1).Methods("GET")
	r.HandleFunc("/tabla2", Tabla2).Methods("GET")
	r.HandleFunc("/tabla3", Tabla3).Methods("GET")
	r.HandleFunc("/file", UploadFile).Methods("POST")
	r.HandleFunc("/Aritmetica", AnalizarExpresionAritmetica).Methods("POST")
	// Vinculacion a puerto y se pasa router para comenzar proceso de escucha
	log.Fatal(http.ListenAndServe(":8001", r))
}

/* Function: getTable1

   Metodo que devuelve un arreglo de simbolos finales que se utilizara en la peticion /page1

   Returns:

      Arreglo de simbolos finales

   See Also:

      <Tabla1>
			<Leer>
			<CrearMapaSimbolos>
			<CrearMapaCorrespondencia>
			<CrearMapaTokens>
*/
//Metodo que devuelve un arreglo de simbolos finales que se utilizara en la peticion /page1
func getTable1() []analizador.FinalSimbol {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura simbolos
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura de simbolos
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
	//Crear mapa de lectura para la correspondencia de simbolos
	tablaCorrespondencia := make(map[string]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de lectura de correspondencia
	analizador.CrearMapaCorrespondencia("TablaCorrespondencia.csv", tablaCorrespondencia)
	//Crear mapa de lectura tokens
	Mapatokens := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de tokens
	analizador.CrearMapaTokens("TablaTokens.csv", Mapatokens)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"

	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := analizador.ContadorLineas(nombreArchivo)

	//Se crea la tabla intermedia que alimenta a la final de tokens
	tablaIntermedia := make([][]analizador.Token, cantidadLineas)

	//Apertura de prog.messi
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		//Busqueda de errores
		fmt.Println(err)
	}
	defer archivo.Close()
	fileScanner := bufio.NewScanner(archivo)
	//Se inicializa la variable que envia la fila que se esta analizando
	i := 0
	//Se lee el archivo linea a linea enviando al metodo leer
	for fileScanner.Scan() {
		_ = analizador.Leer(fileScanner.Text(), i, tablaSimbolos, tablaIntermedia, tablaCorrespondencia, tabla2, Mapatokens)
		i++
	}
	/*Se organiza la tabla intermedia para entregarla como requiere la peticion
	NOTA: tabla intermedia se va a generar en el metodo de leer
	final2: donde se va a alojar todos los simbolos traidos del analisis alojado en tabla intermedia.
	*/
	final2 := []analizador.FinalSimbol{}
	for f, linea := range tablaIntermedia {
		for c, token := range linea {
			var tipos [4]string
			copy(tipos[:], token.TypeToken)
			//Se agrega simbolo en la linea f y posicion c a la tabla final con su respectiva posicion y tipos
			final2 = append(final2, analizador.FinalSimbol{Nombre: token.TokenName, Linea: strconv.Itoa(f), NumSimbFila: strconv.Itoa(c), T1: tipos[0], T2: tipos[1], T3: tipos[2]})
		}

	}
	return final2
}

/*
	 Function: getTable2

	   Metodo que devuelve un arreglo de tokens finales que se utilizara en la peticion /page2

	   Returns:

	      Arreglo de tabla tokens finales

	   See Also:

	      <Tabla2>
				<Leer>
				<CrearMapaSimbolos>
				<CrearMapaCorrespondencia>
				<CrearMapaTokens>
*/
func getTable2() []analizador.TablaTokens {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura simbolos
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura de simbolos
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
	//Crear mapa de lectura para la correspondencia de simbolos
	tablaCorrespondencia := make(map[string]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de lectura de correspondencia
	analizador.CrearMapaCorrespondencia("TablaCorrespondencia.csv", tablaCorrespondencia)
	//Crear mapa de lectura tokens
	Mapatokens := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de tokens
	analizador.CrearMapaTokens("TablaTokens.csv", Mapatokens)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"

	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := analizador.ContadorLineas(nombreArchivo)

	//Se crea la tabla intermedia que alimenta a la final de tokens
	tablaIntermedia := make([][]analizador.Token, cantidadLineas)

	//Apertura de prog.messi
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		//Busqueda de errores
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

/* Function: getTable3

   Metodo que devuelve un arreglo de expresiones aritmeticas que se utilizara en la peticion /page3

   Returns:

      Arreglo de tabla expresiones aritmeticas

   See Also:

      <Tabla3>
			<Leer>
			<CrearMapaSimbolos>
			<CrearMapaCorrespondencia>
			<CrearMapaTokens>
			<RAritmetico>
*/
//Metodo que devuelve un arreglo de expresiones aritmeticas que se utilizara en la peticion /page3
func getTable3() []analizador.AritemticaStruct {
	tabla2 := []analizador.TablaTokens{}
	//Crear mapa de lectura simbolos
	tablaSimbolos := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con los simbolos y crea el mapa de lectura de simbolos
	analizador.CrearMapaSimbolos("TablaSimbolos.csv", tablaSimbolos)
	//Crear mapa de lectura para la correspondencia de simbolos
	tablaCorrespondencia := make(map[string]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de lectura de correspondencia
	analizador.CrearMapaCorrespondencia("TablaCorrespondencia.csv", tablaCorrespondencia)
	//Crear mapa de lectura tokens
	Mapatokens := make(map[string][]string)
	//Se llama a la funcion que lee el el csv con la correspondencia a tipos y crea el mapa de tokens
	analizador.CrearMapaTokens("TablaTokens.csv", Mapatokens)

	//Se ingresa el codigo fuente a analizar
	nombreArchivo := "prog.messi"

	//Se llama a la funcion que cuenta la cantidad de lineas del codigo fuente
	cantidadLineas := analizador.ContadorLineas(nombreArchivo)

	//Se crea la tabla intermedia que alimenta a la final de tokens
	tablaIntermedia := make([][]analizador.Token, cantidadLineas)

	//Apertura de prog.messi
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		//Busqueda de errores
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

	/*Se organiza la tabla intermedia para entregarla como requiere la peticion
	NOTA: tabla intermedia se va a generar en el metodo de leer
	tablaAritmetica: donde se va a alojar todos las expresiones en conjuncion con los datos usados en tablaIntermedia
	*/
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
