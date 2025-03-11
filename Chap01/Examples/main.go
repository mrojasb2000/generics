package main

import (
	"fmt"
)

// GenericFilter filters elements in a slice based on a predicate function.
func GenericFilter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	input := []float64{-5.0, -2.0, 4.0, 8.0}
	result1 := GenericMap(input, func(v float64) float64 {
		return v * v
	})
	fmt.Println(result1)

	greaterThanFive := GenericFilter([]int{4, 6, 5, 2, 20, 1, 7}, func(v int) bool {
		return v > 5
	})
	fmt.Println(greaterThanFive)

	oddNumbers := GenericFilter([]int{4, 6, 5, 2, 20, 1, 7}, func(v int) bool {
		return v%2 == 1
	})
	fmt.Println(oddNumbers)

	lengthGreaterThanThree := GenericFilter([]string{"hello", "or", "the", "maybe"}, func(v string) bool {
		return len(v) > 3
	})
	fmt.Println(lengthGreaterThanThree)
}
