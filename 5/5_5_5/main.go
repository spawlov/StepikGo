package main

import "fmt"

func countDistinctElements(nums []int) int {
	counter := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			counter++
		}
	}

	return counter
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	fmt.Println(countDistinctElements(numbers))
}
