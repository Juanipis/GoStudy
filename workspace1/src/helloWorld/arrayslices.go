package main

import (
	"fmt"
)

func main() {
	//Arrays
	//Se define una cadena con tamaño fijo
	var a [5]int
	//Se ingresan valores a la cadena
	a[0] = 11
	a[1] = 22
	a[2] = 33
	a[3] = 44
	a[4] = 55
	//Se imprime la cadena
	fmt.Println(a)
	//Se imprime la posición de un valor en la cadena
	fmt.Println(a[2])

	//Slices
	//Los slices son arrays que se pueden modificar, se pueden agrandar y se pueden reducir, se pueden truncar, se pueden copiar y se pueden eliminar, se pueden reemplazar
	//Se define un slice
	s := []int{1, 2, 3, 4, 5}
	//Se imprime la cadena
	fmt.Println(s)
	//Se imprime la posición de un valor en la cadena
	fmt.Println(s[2])
	//Se agrega un valor al final del slice
	s = append(s, 6)
	//Se imprime la cadena
	fmt.Println(s)
	//Se toma una porcion de la cadena
	s = s[2:4]
	//Se imprime la cadena
	fmt.Println(s)

	//Para crear slices sin contenido con make
	//Se define un slice
	s1 := make([]int, 3) //Se define un slice de 3 int, el primer parametro es el tipo de dato, el segundo parametro es el tamaño
	//Se imprime la cadena
	fmt.Println(s1)
}
