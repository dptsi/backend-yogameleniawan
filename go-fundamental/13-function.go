package main

import "fmt"

func main() {
	myMessage()
	myName("Yoga Meleniawan Pamungkas")
	myNameAge("Yoga Meleniawan Pamungkas", 23)
	fmt.Println(sum(10, 20))
	fmt.Println(sum2(10, 20))
}

func myMessage() {
	fmt.Println("Hello World")
}

func myName(name string) {
	fmt.Println("Hello my name is ", name)
}

func myNameAge(name string, age int) {
	fmt.Println("Hello my name is ", name, " and I am ", age, " years old")
}

func sum(x int, y int) int {
	return x + y
}

func sum2(x, y int) (result int) {
	result = x + y
	return
}
