package main

import "fmt"

type ejemplo struct {
	texto   string
	arreglo []int
}

func main() {
	arrEjemplo := []ejemplo{}
	ejemplo1 := ejemplo{"hola", []int{1, 2, 3, 4, 5}}
	ejemplo2 := ejemplo{"cosa", []int{6, 1, 99, 1, 89}}
	ejemplo3 := ejemplo{"poll", []int{10, 65, 31, 43, 59}}
	arrEjemplo = append(arrEjemplo, ejemplo1)
	arrEjemplo = append(arrEjemplo, ejemplo2)
	arrEjemplo = append(arrEjemplo, ejemplo3)

	arrEjemplo2 := []ejemplo{}
	ejemplo12 := ejemplo{"hola", []int{1, 2, 3, 4, 5}}
	ejemplo22 := ejemplo{"cosa", []int{6, 1, 99, 1, 89}}
	ejemplo32 := ejemplo{"poll", []int{10, 65, 31, 43, 599}}
	arrEjemplo2 = append(arrEjemplo2, ejemplo12)
	arrEjemplo2 = append(arrEjemplo2, ejemplo22)
	arrEjemplo2 = append(arrEjemplo2, ejemplo32)

	arrDearrEjemplo := [][]ejemplo{arrEjemplo, arrEjemplo2}
	fmt.Println(arrDearrEjemplo[1][2].arreglo[4])
}
