package main

import "fmt"

func main() {
	var i, j string = "Hello", "World"

	// Print() function prints its arguments with their default format.
	fmt.Print(i)
	fmt.Print(j)

	// If we want to print the arguments in new lines, we need to use \n.
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
	fmt.Println(i, j)

	// PrintF()
	// The Printf() function first formats its argument based on the given formatting verb and then prints them.

	// %v is the default format, which is the value in the default format:
	// %T is the type of the value:

	var k string = "Hello"
	var l int = 15

	fmt.Printf("k has value: %v and type: %T\n", k, l)
	fmt.Printf("l has value: %v and type: %T", k, l)
}
