package main

import "fmt"

func minValue(array []int) int {
	result := array[0]
	for _, value := range array {
		if value < result {
			result = value
		}
	}
	return result
}

func decreaseValues(array []int, number int) []int {
	for i := 0; i < len(array); i++ {
		array[i] = array[i] - number
	}
	return array
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	result := decreaseValues(numbers, minValue(numbers))
	for _, value := range result {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
}
