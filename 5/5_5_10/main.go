package main

import "fmt"

func main() {
	var n, m int

	_, _ = fmt.Scan(&n, &m)

	photoColor := false
	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, m)
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&matrix[i][j])
			if matrix[i][j] == "C" || matrix[i][j] == "M" || matrix[i][j] == "Y" {
				photoColor = true
			}
		}
	}
	
	if photoColor {
		fmt.Println("#Color")
	} else {
		fmt.Println("#Black&White")
	}
}
