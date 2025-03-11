package main

import "fmt"

func MyFilter(input []float64, f func(float64) bool) []float64 {
	result := make([]float64, 0)
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	input := []float64{17.3, 11.1, 9.9, 4.3, 12.6}
	result := MyFilter(input, func(v float64) bool {
		return v <= 10.0
	})
	fmt.Println(result)

}
