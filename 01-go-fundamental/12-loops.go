package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop with condition
	var j = 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// for loop with range
	var k = []int{1, 2, 3, 4, 5}
	for l, m := range k {
		fmt.Println(l, m)
	}

	// nested for loop
	for n := 0; n < 5; n++ {
		for o := 0; o < 5; o++ {
			fmt.Println(n, o)
		}
	}

	// infinite for loop
	for {
		fmt.Println("Infinite loop")
		break
	}
}
