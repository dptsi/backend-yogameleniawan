package main

import (
	"fmt"
	"math"
)

// /* define an interface */
// type interface_name interface {
// 	method_name1 [return_type]
// 	method_name2 [return_type]
// 	method_name3 [return_type]
// 	...
// 	method_namen [return_type]
//  }

//  /* define a struct */
//  type struct_name struct {
// 	/* variables */
//  }

//  /* implement interface methods*/
//  func (struct_name_variable struct_name) method_name1() [return_type] {
// 	/* method implementation */
//  }
//  ...
//  func (struct_name_variable struct_name) method_namen() [return_type] {
// 	/* method implementation */
//  }

/* define an interface */
type Shape interface {
	area() float64
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
	width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
	return shape.area()
}

type Bentuk interface {
	luas() float64
}

type Persegi struct {
	sisi float64
}

type Segitiga struct {
	alas, tinggi float64
}

func (persegi Persegi) luas() float64 {
	return persegi.sisi * persegi.sisi
}

func (segitiga Segitiga) luas() float64 {
	return 0.5 * segitiga.alas * segitiga.tinggi
}

func getLuas(bentuk Bentuk) float64 {
	return bentuk.luas()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	persegi := Persegi{sisi: 5}

	fmt.Printf("Luas persegi: %f\n", getLuas(persegi))

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
