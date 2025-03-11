package main

import "fmt"

// MyMap is a function produces an output slice
// containing the results of applying the function "f"
// to each element of the input slice.
func MyMap(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	slice := []int{1, 5, 2, 7, 4}
	result := MyMap(slice, func(v int) int {
		return v * v
	})
	fmt.Println(result)
}
