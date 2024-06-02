package main

import "fmt"

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	minValue := numbers[0]
	minIndex := 0
	for index, value := range numbers {
		if value < minValue {
			minValue = value
			minIndex = index
		}
	}
	fmt.Println(minIndex)
}
