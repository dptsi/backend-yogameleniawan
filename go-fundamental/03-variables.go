package main

import (
	"fmt"
)

func main() {
	var student1 string = "Yoga" //type is string
	var student2 = "Pamungkas"   //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Variable Declaration Without Initial Value

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//   Value Assignment After Declaration

	var student3 string
	student3 = "Yoga"
	fmt.Println(student3)

	//   There are some small differences between the var var :=:

	// var :
	// Can be used inside and outside of functions
	// Variable declaration and value assignment can be done separately

	// := :
	// Can only be used inside functions
	// Variable declaration and value assignment cannot be done separately (must be done in the same line)

	//  Multiple Variable Declaration
	multipleVariable()
	multipleVariableDifferentTypes()
	variableBlock()
}

func multipleVariable() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//   Multiple Variable Declaration with Different Types

func multipleVariableDifferentTypes() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func variableBlock() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
