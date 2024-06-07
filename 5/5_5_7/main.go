package main

import "fmt"

func createArray(n int) []int {
	array := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&array[i])
	}
	return array
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := createArray(n)

	for idx := range numbers {
		for i := idx + 1; i < n; i++ {
			if numbers[idx]+numbers[i] == 0 {
				fmt.Println(idx, i)
				break
			}
		}
	}
}
