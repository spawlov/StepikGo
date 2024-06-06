package main

import "fmt"

func createMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func sumNumbers(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func main() {
	var n, m int

	_, _ = fmt.Scan(&n, &m)

	matrix := createMatrix(n, m)

	sumIndex := 0
	sumResult := 0
	for index, array := range matrix {
		if sumResult < sumNumbers(array) {
			sumIndex = index
			sumResult = sumNumbers(array)
		}
	}
	fmt.Print(sumResult, "\n", sumIndex, "\n")
}
