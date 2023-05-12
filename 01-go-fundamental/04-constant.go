package main

import (
	"fmt"
)

const PI = 3.14

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println(PI)

	const A = 1
	A = 2 // Constants: Unchangeable and Read-only
	// it doesn't work to assign/change value to constant
	fmt.Println(A)

	multipleConstDeclaration()
}

func multipleConstDeclaration() {

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
