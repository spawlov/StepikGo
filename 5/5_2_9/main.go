package main

import "fmt"

func firstMin(array []int) int {
	minIdx := 0
	minVal := array[minIdx]
	for i := 0; i < len(array); i++ {
		if array[i] < minVal {
			minVal = array[i]
			minIdx = i
		}
	}
	return minIdx
}

func lastMax(array []int) int {
	maxIdx := 0
	maxVal := array[maxIdx]
	for i := 0; i < len(array); i++ {
		if array[i] >= maxVal {
			maxVal = array[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	numbers[lastMax(numbers)], numbers[firstMin(numbers)] = numbers[firstMin(numbers)], numbers[lastMax(numbers)]
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Print("\n")
}
