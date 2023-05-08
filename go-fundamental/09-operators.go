package main

import "fmt"

func main() {
	// Arithmetic Operators ( +, -, *, /, %, ++, --)
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	// Relational Operators ( ==, !=, >, <, >=, <=)
	var d = a > b
	fmt.Println(d)

	// Logical Operators ( &&, ||, !)
	var e = a > b && a < b
	fmt.Println(e)

	// Bitwise Operators ( &, |, ^, <<, >>)
	var f = a & b
	fmt.Println(f)

	// Assignment Operators ( =, +=, -=, *=, /=, %=, <<=, >>=, &=, ^=, |=)
	var g = a
	g += 10
	fmt.Println(g)

	// Misc Operators ( &, *, <-, &, <-)
	var h = &a
	fmt.Println(h)
}
