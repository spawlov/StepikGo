package main

import "fmt"

func main() {
	var n, idxMin, idxMax int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	minValue := numbers[0]
	maxValue := numbers[0]
	for index, value := range numbers {
		if value > maxValue {
			idxMax = index
			maxValue = value
		} else if value < minValue {
			idxMin = index
			minValue = value
		}
	}
	fmt.Println(idxMax - idxMin)
}
