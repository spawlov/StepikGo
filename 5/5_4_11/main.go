package main

import "fmt"

func createMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func isMatrixSimmetric(matrix [][]int) string {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	matrix := createMatrix(n)

	fmt.Println(isMatrixSimmetric(matrix))
}
