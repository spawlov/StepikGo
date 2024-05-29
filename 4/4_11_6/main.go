package main

import "fmt"

func main() {
	var N, result int

	_, _ = fmt.Scan(&N)

	numbers := make([]int, N)
	minimum := int(^uint(0) >> 1)

	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&numbers[i])
		if numbers[i] < minimum {
			minimum = numbers[i]
		}
	}

	for _, number := range numbers {
		if number == minimum {
			result++
		}
	}
	fmt.Println(result)
}
