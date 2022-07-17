package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Lectura completa de un archivo
	//Se define la ruta del archivo
	nombreArchivo := "empleados.txt"
	//Se lee el archivo, bytes es el tipo de dato que devuelve el método ReadFile, el cual devuelve un slice de bytes, err es el tipo de dato que devuelve el método ReadFile, el cual devuelve un error
	bytesLeidos, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Println(err)
	}
	//Se convierte el slice de bytes a string
	contenido := string(bytesLeidos)
	fmt.Println(contenido)

	fmt.Println("----------------------------------------------------")
	//Lectura linea a linea de un archivo
	//Se define la ruta del archivo
	nombreArchivo = "empleados.txt"
	//Se abre el archivo
	archivo, err := os.Open(nombreArchivo)

	//Se comprueba si hay un error
	if err != nil {
		fmt.Println(err)
	}

	//Se declara el cerrado del archivo al finalizar la ejecución
	defer archivo.Close()

	//Se crea un buffer de lectura
	fileScanner := bufio.NewScanner(archivo)

	//Se lee el archivo linea a linea
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	//Se comprueba si hay un error
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}

}
