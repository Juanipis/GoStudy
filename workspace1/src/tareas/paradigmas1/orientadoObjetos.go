package main

import "fmt"

type datosUsuarios struct {
	usuarios []dato
}

type dato struct {
	nombre string
	edad   int
}

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

func (f *datosUsuarios) leerDato() {
	f.usuarios = append(f.usuarios, dato{leerNombres(), leerEdad()})
}

func (f *datosUsuarios) sumaEdades() int {
	var suma int
	for _, dato := range f.usuarios {
		suma += dato.edad
	}
	return suma
}

func (f *datosUsuarios) promedioEdades() int {
	return f.sumaEdades() / len(f.usuarios)
}

func (f *datosUsuarios) mayorEdad() dato {
	mayor := f.usuarios[0]
	for _, dato := range f.usuarios {
		if dato.edad > mayor.edad {
			mayor = dato
		}
	}
	return mayor
}

func (f *datosUsuarios) menorEdad() dato {
	menor := f.usuarios[0]
	for _, dato := range f.usuarios {
		if dato.edad < menor.edad {
			menor = dato
		}
	}
	return menor
}

func main() {
	datosGuardados := datosUsuarios{}
	datosGuardados.leerDato()
	datosGuardados.leerDato()
	datosGuardados.leerDato()

	fmt.Printf("El promedio de las edades es %v\n", datosGuardados.promedioEdades())
	fmt.Printf("La suma de las edades es %v\n", datosGuardados.sumaEdades())
	fmt.Printf("La persona de mayor edad es: %v\n", datosGuardados.mayorEdad())
	fmt.Printf("La persona de menor edad es: %v\n", datosGuardados.menorEdad())
}
