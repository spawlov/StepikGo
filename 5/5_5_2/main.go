package main

import "fmt"

func evenValues(array []int) []int {
	var result []int
	for _, value := range array {
		if value%2 == 0 {
			result = append(result, value)
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

	for _, value := range evenValues(numbers) {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
	
	fmt.Println(len(evenValues(numbers)))
}
