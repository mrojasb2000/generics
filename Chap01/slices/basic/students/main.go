package main

import "fmt"

type student struct {
	Name string
	ID   int
	age  float64
}

// The function addStudent takes a slice of string representing the current collection of
// students as the first parameter and a string representing a new student to be added to
// the collection as the second parameter.
func addStudent(students []string, student string) []string {
	return append(students, student)
}

// The function addStudentID takes a slice of int representing the current collection of
// students ID's as the first parameter and a int representing a new student ID to be added to
// the collection as the second parameter.
func addStudentID(students []int, student int) []int {
	return append(students, student)
}

// The function addStudentStruct takes a slice of student type representing the current collection of
// students as the first parameter and a student representing a new student instance to be added to
// the collection as the second parameter.
func addStudentStruct(students []student, student student) []student {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudentID(students1, 155)
	result1 = addStudentID(result1, 112)
	result1 = addStudentID(result1, 120)
	fmt.Println(result1)

	students2 := []student{}
	result2 := addStudentStruct(students2, student{"John", 213, 17.5})
	result2 = addStudentStruct(result2, student{"James", 111, 18.75})
	result2 = addStudentStruct(result2, student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
