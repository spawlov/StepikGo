package main

import "fmt"

func createMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
			}
		}
	}
	return matrix
}

func main() {
	var n, m int

	_, _ = fmt.Scan(&n, &m)

	matrix := createMatrix(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
