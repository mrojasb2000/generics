package main

import "fmt"

func GenericFilter[T any](input []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	result := GenericFilter(input, func(v float64) bool {
		return v <= 10.0
	})
	fmt.Println(result)
}
