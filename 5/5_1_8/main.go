package main

import "fmt"

func main() {
	var N int

	_, _ = fmt.Scan(&N)

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	result := 0
	for i := 1; i < N; i++ {
		if numbers[i-1] < numbers[i] {
			result++
		}
	}
	fmt.Println(result)
}
