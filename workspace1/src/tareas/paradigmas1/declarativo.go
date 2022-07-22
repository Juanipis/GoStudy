package main

import "fmt"

//Paso 0: Estructura para almacenar datos de entrada
type datos struct {
	nombre string
	edad   int
}

//Paso 1: conseguir leer datos de entrada
func leerEdad() int {
	var edad int
	fmt.Printf("Ingrese su edad: ")
	fmt.Scan(&edad)
	return edad
}
func leerNombres() string {
	var nombre string
	fmt.Printf("Ingrese su nombre: ")
	fmt.Scan(&nombre)
	return nombre
}

//Paso 2: Crear estructura para almacenar datos de entrada
func nuevoDatos() datos {
	dato := datos{leerNombres(), leerEdad()}
	return dato
}

//Paso 3: Crear funcion para obtener la suma de las edades de un arreglo de datos
func sumaEdades(datos []datos) int {
	var suma int
	for _, dato := range datos {
		suma += dato.edad
	}
	return suma
}

//Paso 4: Crear funcion para obtener la edad promedio de un arreglo de datos
func edadPromedio(datos []datos) int {
	return sumaEdades(datos) / len(datos)
}

//Paso 5: Obtener la persona con la mayor edad
func mayorEdad(datos []datos) datos {
	mayor := datos[0]
	for _, dato := range datos {
		if dato.edad > mayor.edad {
			mayor = dato
		}
	}
	return mayor
}

//Paso 6: Crear funcion para obtener la persona con la menor edad
func menorEdad(datos []datos) datos {
	menor := datos[0]
	for _, dato := range datos {
		if dato.edad < menor.edad {
			menor = dato
		}
	}
	return menor
}

//Paso 7: leer de manera repetitiva los datos de entrada
func leerDatos() []datos {
	var datos []datos
	for {
		dato := nuevoDatos()
		datos = append(datos, dato)
		fmt.Printf("Desea ingresar otro dato? (s/n)")
		var respuesta string
		fmt.Scan(&respuesta)
		if respuesta == "n" {
			break
		}
	}
	return datos
}

func main() {
	//1. Leer los datos
	datosUsuarios := leerDatos()
	//2. Imprimir la edad promedio
	fmt.Printf("La edad promedio de los usuarios es: %d\n", edadPromedio(datosUsuarios))
	//3. Imprimir la suma de edades
	fmt.Printf("La suma de las edades de los usarios es %d\n", sumaEdades(datosUsuarios))
	//4. Imprimir la persona de mayor edad
	fmt.Printf("La persona de mayor edad es: %v\n", mayorEdad(datosUsuarios))
	//5. Imprimir la persona de menor edad
	fmt.Printf("La persona de menor edad es: %v\n", menorEdad(datosUsuarios))

}
