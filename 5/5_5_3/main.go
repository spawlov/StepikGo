package main

import "fmt"

func lessThanNext(array []int) int {
	result := 0
	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			result++
		}
	}

	return result
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	fmt.Println(lessThanNext(numbers))
}
