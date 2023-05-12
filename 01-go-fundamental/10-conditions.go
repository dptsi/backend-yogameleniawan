package main

import "fmt"

func main() {
	// if else
	var a = 10
	var b = 20
	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	// if else if
	var c = 30
	if a > b {
		fmt.Println("a is greater than b")
	} else if b > c {
		fmt.Println("b is greater than c")
	} else {
		fmt.Println("c is greater than a and b")
	}

	// switch case
	var d = 40
	switch d {
	case 10:
		fmt.Println("d is 10")
	case 20:
		fmt.Println("d is 20")
	case 30:
		fmt.Println("d is 30")
	default:
		fmt.Println("d is not 10, 20, or 30")
	}

	// switch case with multiple values
	var e = 50
	switch e {
	case 10, 20, 30:
		fmt.Println("e is either 10, 20, or 30")
	default:
		fmt.Println("e is not 10, 20, or 30")
	}

	// switch case with condition
	var f = 60
	switch {
	case f > 10:
		fmt.Println("f is greater than 10")
	case f > 20:
		fmt.Println("f is greater than 20")
	case f > 30:
		fmt.Println("f is greater than 30")
	default:
		fmt.Println("f is not greater than 10, 20, or 30")
	}
}
