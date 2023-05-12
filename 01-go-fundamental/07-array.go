package main

import "fmt"

func main() {
	// var keyword

	// var array_name = [length]datatype{values} // here length is defined
	// var array_name = [...]datatype{values} // here length is inferred

	var array_name = [5]int{1, 2, 3, 4, 5}
	var array_name2 = [...]int{1, 2, 3, 4, 5}

	fmt.Println(array_name)
	fmt.Println(array_name2)

	// := sign

	// array_name := [length]datatype{values} // here length is defined
	// array_name := [...]datatype{values} // here length is inferred

	array_name3 := [5]int{1, 2, 3, 4, 5}
	array_name4 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(array_name3)
	fmt.Println(array_name4)

	// access array elements

	buah := [5]string{"apel", "jeruk", "mangga", "melon", "semangka"}
	fmt.Println(buah[0]) // apel
	fmt.Println(buah[1]) // jeruk
}
