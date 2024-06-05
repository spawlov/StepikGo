package main

import "fmt"

func hasDuplicate(array []int) string {
	seen := make(map[int]bool)

	for _, number := range array {
		if seen[number] {
			return "YES"
		}
		seen[number] = true
	}
	return "NO"
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	fmt.Println(hasDuplicate(numbers))
}
