package main

import "fmt"

//Paso 0: Estructura para almacenar datos de entrada
type datos struct {
	nombre string
	edad   int
}

func main() {
	//1. Obtener datos de entrada
	var datosUsuarios []datos
	for {
		//Obtener el dato
		var nombre string
		var edad int
		fmt.Printf("Ingrese su nombre: ")
		fmt.Scan(&nombre)
		fmt.Printf("Ingrese su edad: ")
		fmt.Scan(&edad)
		dato1 := datos{nombre, edad}
		datosUsuarios = append(datosUsuarios, dato1)
		//Continuar ingresando datos
		fmt.Printf("Desea ingresar otro dato? (s/n)")
		var respuesta string
		fmt.Scan(&respuesta)
		if respuesta == "n" {
			break
		}
	}

	//2. Imprimir salidas pedidas
	//2.1. Promedio edades edades
	var suma int
	for _, dato := range datosUsuarios {
		suma += dato.edad
	}
	//Imprimir promedio edades
	fmt.Printf("La edad promedio de los usuarios es: %d\n", suma/len(datosUsuarios))
	//Imprimir suma edades
	fmt.Printf("La suma de las edades de los usarios es %d\n", suma)

	//2.2 Imprimir la persona de mayor edad
	mayor := datosUsuarios[0]
	for _, dato := range datosUsuarios {
		if dato.edad > mayor.edad {
			mayor = dato
		}
	}
	fmt.Printf("La persona de mayor edad es: %v\n", mayor)

	//2.3 Imprimir la persona de menor edad
	menor := datosUsuarios[0]
	for _, dato := range datosUsuarios {
		if dato.edad < menor.edad {
			menor = dato
		}
	}
	fmt.Printf("La persona de menor edad es: %v\n", menor)

}
