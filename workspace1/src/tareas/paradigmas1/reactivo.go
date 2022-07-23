package main

import (
	"fmt"
)

type datoUsuarios struct {
	nombre string
	edad   int
}

func lecturaDatos(chNombre chan<- string, chEdad chan<- int) {
	for {
		var nombre string
		var edad int
		fmt.Printf("Ingrese su nombre: ")
		fmt.Scan(&nombre)
		fmt.Printf("Ingrese su edad: ")
		fmt.Scan(&edad)

		fmt.Printf("Desea ingresar otro dato? (s/n)")
		var respuesta string
		fmt.Scan(&respuesta)
		if respuesta == "n" {
			break
		} else {
			chNombre <- nombre
			chEdad <- edad
		}
	}
}

func main() {
	datos := []datoUsuarios{}
	chNombre := make(chan string)
	chEdad := make(chan int)

	go lecturaDatos(chNombre, chEdad)

	//Ciclo de espera de datos
	for {
		datos = append(datos, datoUsuarios{<-chNombre, <-chEdad})
		//Imprimamos los datos actuales
		var suma int
		for _, dato := range datos {
			suma += dato.edad
		}
		fmt.Printf("La edad promedio de los usuarios es: %d\n", suma/len(datos))
		//Imprimir suma edades
		fmt.Printf("La suma de las edades de los usarios es %d\n", suma)

		//Imprimir mayor edad
		mayor := datos[0]
		for _, dato := range datos {
			if dato.edad > mayor.edad {
				mayor = dato
			}
		}
		fmt.Printf("La persona de mayor edad es: %v\n", mayor)
		//Imprimir menor edad
		menor := datos[0]
		for _, dato := range datos {
			if dato.edad < menor.edad {
				menor = dato
			}
		}
		fmt.Printf("La persona de menor edad es: %v\n", menor)
	}

}
