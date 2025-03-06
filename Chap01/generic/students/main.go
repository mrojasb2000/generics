package main

import "fmt"

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
}
