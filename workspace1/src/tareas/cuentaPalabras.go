package main

import (
	"fmt"
	"strconv"
)

func contadorPalabrasReservadas(cadena string) (map[int]string, map[int]string, map[int]string) {
	mapaWhile := isWord(cadena, "WHILE")
	mapaIf := isWord(cadena, "IF")
	mapaFor := isWord(cadena, "FOR")
	return mapaWhile, mapaIf, mapaFor
}

func isWord(cadena string, word string) map[int]string {
	//Buscar WHILE
	var mapa map[int]string
	mapa = make(map[int]string)
	ini := 0
	count := 0
	for ini+len(word) <= len(cadena) {
		if cadena[ini:ini+len(word)] == word {
			mapa[count] = strconv.Itoa(ini) + "," + strconv.Itoa(ini+len(word))
			count++
			ini = ini + len(word)
		} else {
			ini++
		}
	}
	return mapa

}

func main() {
	cadena := " WHILE WHILE FORIF"
	mapaWhile, mapaIf, mapaFor := contadorPalabrasReservadas(cadena)
	fmt.Println(mapaWhile)
	fmt.Println(mapaIf)
	fmt.Println(mapaFor)

}
