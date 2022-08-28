package analizador

import (
	"fmt"

	"golang.org/x/exp/slices"
)

var constanteVariable []string
var operadores []string
var almFinal []TablaTokens

func RAritmetica(tabla2 []TablaTokens) {
	constanteVariable = []string{"IdentificadorVariable", "IdentificadorConstante"}
	operadores = []string{"+", "-", "*", "/", "%", "^", "¬", ":^"}
	almExp := []TablaTokens{}
	//(a+b)-c+

	//a
	varConst1(tabla2, 0, almExp)

	//+
	//op2()

	//b
	//varConst2()
	fmt.Println(almFinal)

}

func varConst1(tabla2 []TablaTokens, indice int, almExp []TablaTokens) {
	if indice < len(tabla2) && slices.Contains(constanteVariable, tabla2[indice].Token) {
		almExp = append(almExp, tabla2[indice])
		//almFinal = append(almFinal, tabla2[indice])
		op2(tabla2, indice+1, almExp)
	}
	if indice < len(tabla2) && tabla2[indice].LexemaGenerador == "(" {
		almExp = append(almExp, tabla2[indice])
		//almFinal = append(almFinal, tabla2[indice])
		varConst1(tabla2, indice+1, almExp)
		if len(almFinal) < len(tabla2) && tabla2[len(almFinal)].LexemaGenerador == ")" {
			almFinal = append(almFinal, tabla2[len(almFinal)])
			op2(tabla2, len(almFinal), almExp)
		}
	} else if indice+1 < len(tabla2) {
		almExp = make([]TablaTokens, 0)
		varConst1(tabla2, indice+1, almExp)
	}
}

func op2(tabla2 []TablaTokens, indice int, almExp []TablaTokens) {
	if indice < len(tabla2) && slices.Contains(operadores, tabla2[indice].LexemaGenerador) {
		almExp = append(almExp, tabla2[indice])
		//almFinal = append(almFinal, tabla2[indice])
		varConst2(tabla2, indice+1, almExp)
	} else if indice+1 < len(tabla2) {
		almExp = make([]TablaTokens, 0)
		varConst1(tabla2, indice+1, almExp)
	}
}

func varConst2(tabla2 []TablaTokens, indice int, almExp []TablaTokens) {
	if indice < len(tabla2) && slices.Contains(constanteVariable, tabla2[indice].Token) {
		almExp = append(almExp, tabla2[indice])
		almFinal = append(almFinal, almExp...)
	}
	if indice < len(tabla2) && tabla2[indice].LexemaGenerador == "(" {
		almExp = append(almExp, tabla2[indice])
		//almFinal = append(almFinal, tabla2[indice])
		varConst1(tabla2, indice+1, almExp)
		if len(almFinal) < len(tabla2) && tabla2[len(almFinal)].LexemaGenerador == ")" {
			almFinal = append(almFinal, tabla2[len(almFinal)])
			if len(almFinal) < len(tabla2) {
				almExp = make([]TablaTokens, 0)
				op2(tabla2, len(almFinal), almExp)
			}
		}
	}

}
