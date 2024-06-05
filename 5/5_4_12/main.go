package main

import "fmt"

func createMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}
	return matrix
}

func main() {
	var n, m int

	_, _ = fmt.Scan(&n, &m)

	matrix:=createMatrix(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
