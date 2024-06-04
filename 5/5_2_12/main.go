package main

import "fmt"

func isPalindrome(array []int) string {
	left := 0
	right := len(array) - 1

	for left < right {
		if array[left] != array[right] {
			return "NO"
		}
		left++
		right--
	}
	return "YES"
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	fmt.Println(isPalindrome(numbers))
}
