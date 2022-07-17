package main

import "fmt"

func main() {
	//Imprimir en pantalla
	fmt.Println("fmt para imprimir en pantalla")
	//Ingresar datos por consola
	var edad int
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&edad)
	fmt.Println("Su edad es: ", edad)
}
