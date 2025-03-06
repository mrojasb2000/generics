package main

import "fmt"

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
}
