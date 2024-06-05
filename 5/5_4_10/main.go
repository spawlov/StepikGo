package main

import "fmt"

func createMatrix(n int) [][]int {
	array := make([][]int, n)

	for i := 0; i < n; i++ {
		array[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				array[i][j] = 1
			} else if i+j < n-1 {
				array[i][j] = 0
			} else {
				array[i][j] = 2
			}
		}
	}
	return array
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	matrix := createMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
