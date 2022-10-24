package estructurasdatos

// Class: PilaString
// Clase que representa una pila de strings.
type PilaString struct {
	// Variable: Pila
	// La pila de tipo string.
	Pila stack[string]
}

// Class: stack
// Clase que representa una pila generica.
type stack[T any] struct {
	/*
		Function: Push
		Funccion que agrega un elemento a la pila.
		Parameters:
			T - Elemento a agregar a la pila.
	*/
	Push func(T)

	/*
		Function: Pop
		Funccion que elimina el ultimo elemento de la pila.
		Returns:
			T - Elemento eliminado de la pila.
	*/
	Pop func() T
	/*
		Function: Length
		Funccion que retorna el tamaño de la pila.
		Returns:
			int - Tamaño de la pila.
	*/
	Length func() int
	/*
		Function: Head
		Funccion que retorna el ultimo elemento de la pila.
		Returns:
			T - Ultimo elemento de la pila.
	*/
	Head func() T
}

func Stack[T any]() stack[T] {
	slice := make([]T, 0)
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
		Head: func() T {
			return slice[len(slice)-1]
		},
	}
}
