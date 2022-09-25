package analizador

// Class: Token
// Generacion de estructura para manejo de simbolo\nSu existencia es necesaria para la funcion principal de lectura
type Token struct {
	// Variable: TokenName
	// El nombre del token
	TokenName string
	// Variable: TypeToken
	// El arreglo de los tipos del token
	TypeToken []string
}

// Class: ExprAritmetica
// Generacion de estructura para manejo de expresiones aritmeticas\nSu existencia es necesaria para recibir los datos en formato json y poder ser leidos por el backend
type ExprAritmetica struct {
	// Variable: Exp
	// La expresion aritmetica
	Exp string `json:"exp"`
}

// Class: ResultAritmetica
// Generacion de estructura para manejo de expresiones aritmeticas\nSu existencia es necesaria para enviar los datos en formato json y poder ser leidos por el frontend
type ResultAritmetica struct {
	// Variable: Result
	// El resultado de la expresion aritmetica, puede ser True o False, en el front se traduce a ACEPTADO o RECHAZADO
	Result bool `json:"result"`
	// Variable: Log
	// El log de la expresion aritmetica, contiene los errores que se hayan encontrado y los match que se hayan hecho
	Log string `json:"log"`
}

// Class: TablaTokens
// Generacion de estructura para manejo de los tokens del codigo fuente
type TablaTokens struct {
	// Variable: Token
	// El nombre del Token, este es definido por el desarrollador
	Token string
	// Variable: IdToken
	// El id del token, este es único segun el token
	IdToken string
	// Variable: Lexema
	// El lexema generado correspondiente al token ()
	LexemaGenerador string
}

// Class: FinalSimbol
// Generacion de estructura para manejo de la tabla de simbolos que surge del analisis.\nSe compone de aquellos datos asociados a las columnas (Tabla1) que van a ser visibles en el programa
type FinalSimbol struct {
	// Variable: Nombre
	// El nombre del simbolo
	Nombre string `json:"name"`
	// Variable: Linea
	// La linea en la que se encuentra el simbolo
	Linea string `json:"line"`
	// Variable: NumSimbFila
	// La columna en la que se encuentra el simbolo
	NumSimbFila string `json:"numSimbFila"`
	// Variable: T1
	// El tipo 1 del simbolo
	T1 string `json:"t1"`
	// Variable: T2
	// El tipo 2 del simbolo
	T2 string `json:"t2"`
	// Variable: T3
	// El tipo 3 del simbolo
	T3 string `json:"t3"`
}

/*
Class: AritemticaStruct
Estructura que guarda toda la informacion asociada a la expresion aritmetica y que debera ser almacenada
*/
type AritemticaStruct struct {
	// Variable: Expresion
	// Expresion aritmetica
	Expresion string `json:"exp"`
	// Variable: Linea
	// Linea del codigo en la que se encuentra la expresion aritmetica
	Linea string `json:"linea"`
	// Variable: SimInicial
	// Ubicacion del simbolo inicial de la expresion aritmetica en la linea de codigo correspondiente
	SimInicial string `json:"simInicio"`
	// Variable: SimInicial
	// Ubicacion del simbolo final de la expresion aritmetica en la linea de codigo correspondiente
	SimFinal string `json:"simFinal"`
}
