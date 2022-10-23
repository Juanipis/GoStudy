package main

import (
	"fmt"
)

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
	Head   func() T
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
func main() {
	stack := Stack[string]()
	stack.Push("this")
	stack.Push("is")
	stack.Push("a")
	stack.Push("test")
	fmt.Println(stack.Length())
	fmt.Println(stack.Head())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
}
