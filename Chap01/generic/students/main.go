package main

import "fmt"

// data type restricted by Stringer interfaces
type Stringer = interface {
	String() string
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name String
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

// The function addStudent takes a slice of T type (data type restricted by Stringer interfaces)
// representing the current collection of T type as the first parameter and a T type representing
// a new instance to be added to the collection as the second parameter.
func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

func main() {
	students := []String{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)

	students1 := []Integer{}
	result1 := addStudent(students1, 45)
	result1 = addStudent(result1, 64)
	result1 = addStudent(result1, 78)
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent(students2, Student{Name: "John", ID: 213, Age: 17.5})
	result2 = addStudent(result2, Student{Name: "James", ID: 111, Age: 18.75})
	result2 = addStudent(result2, Student{Name: "Marsha", ID: 110, Age: 16.25})
	fmt.Println(result2)

}
