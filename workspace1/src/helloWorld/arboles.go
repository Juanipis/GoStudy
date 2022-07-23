package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	tree := &Tree{
		&Tree{
			&Tree{nil, 1, nil},
			2,
			&Tree{nil, 3, nil},
		},
		4,
		&Tree{
			&Tree{nil, 5, nil},
			6,
			&Tree{nil, 7, nil},
		},
	}
	fmt.Println(tree.Left.Value)
}
