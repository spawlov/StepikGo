package main

import "fmt"

func mathSign(arr []int, n int) string {
	for i := 1; i < n; i++ {
		if arr[i-1] < 0 && arr[i] < 0 {
			return "YES"
		} else if arr[i-1] > 0 && arr[i] > 0 {
			return "YES"

		}
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
	fmt.Println(mathSign(numbers, n))
}
