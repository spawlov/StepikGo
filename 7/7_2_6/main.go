package main

import "fmt"

func main() {
	maxValueOne := findMaxValue(createArray())
	maxValueTwo := findMaxValue(createArray())
	fmt.Println(maxValueOne * maxValueTwo)
}

func createArray() []int {
	var n int
	_, _ = fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&array[i])
	}
	return array
}

func findMaxValue(array []int) int {
	result := array[0]
	for _, value := range array {
		if value > result {
			result = value
		}
	}
	return result
}
