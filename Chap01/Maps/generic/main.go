package main

import "fmt"

func GenericMap[T1, T2 any](input []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	slice := []int{1, 5, 2, 7, 4}
	result := GenericMap(slice, func(v int) int {
		return v * v
	})
	fmt.Println(result)
}
