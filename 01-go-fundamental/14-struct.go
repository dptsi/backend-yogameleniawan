package main

import "fmt"

// type struct_name struct {
// 	member1 datatype;
// 	member2 datatype;
// 	member3 datatype;
// 	...
// }

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	var student1 Student
	student1.name = "Yoga Meleniawan Pamungkas"
	student1.age = 23
	student1.grade = 3

	var student2 = Student{"Pamungkas", 23, 3}

	fmt.Println("Student 1 name: ", student1.name, ", age: ", student1.age, ", grade: ", student1.grade)
	fmt.Println("Student 2 name: ", student2.name, ", age: ", student2.age, ", grade: ", student2.grade)

	// Print Student 1 by calling a function printPerson
	printPerson(student1)
}

func printPerson(stud Student) {
	fmt.Println("Name: ", stud.name)
	fmt.Println("Age: ", stud.age)
	fmt.Println("Grade: ", stud.grade)
}
