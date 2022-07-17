package main

import "fmt"

func main() {
	//Estructuras basicas de control
	//Variables
	num1 := 11

	//If, else if, else
	if num1 < 11 {
		fmt.Println("num1 es menor que 11")
	} else if num1 == 11 {
		fmt.Println("num1 es igual que 11")
	} else {
		fmt.Println("num1 es mayor que 11")
	}

	//For
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("i = ", i)
	}

	//Switch
	num2 := 3
	switch num2 {
	case 1:
		fmt.Println("num2 es 1")
	case 2:
		fmt.Println("num2 es 2")
	case 3:
		fmt.Println("num2 es 3")
	case 4:
		fmt.Println("num2 es 4")
	case 5:
		fmt.Println("num2 es 5")
	}

	//While, en GO no existe el while, se usa for para iterar mientras una condicion sea verdadera
	num3 := 0
	suma := 0
	for num3 < 12 {
		suma += num3
		num3++
	}
	fmt.Println("La suma de los numeros del 0 al 11 es: ", suma)

	//break, rompe el ciclo
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("i = ", i)
	}
	fmt.Println("------")

	//continue, continua con el ciclo
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("i = ", i)
	}

	//Nombrando ciclos
loop:
	for {
		for {
			fmt.Println("Rompiendo el bucle externo")
			break loop
		}
		fmt.Println("Esto nunca se imprimirá en pantalla")
	}

	//Defer, se ejecuta al finalizar el programa
	defer fmt.Println("Esto se ejecuta al finalizar el programa")
	fmt.Println("Esto se ejecuta antes del finalizar el programa")
}
