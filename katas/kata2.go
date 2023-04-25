package main

import "fmt"

func tribonacci(signature [3]float64, n int) []float64 {
	result := make([]float64, n)
	copy(result, signature[:])
	for i := 3; i < n; i++ {
		result[i] = result[i-1] + result[i-2] + result[i-3]
	}
	return result
}

func main() {
	result1 := tribonacci([3]float64{1, 1, 1}, 9)  // result1 will be [1, 1, 1, 3, 5, 9, 17, 31, 57]
	result2 := tribonacci([3]float64{0, 0, 1}, 10) // result2 will be [0, 0, 1, 1, 2, 4, 7, 13, 24, 44]
	fmt.Println(result1)
	fmt.Println(result2)
}