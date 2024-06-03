package main

import "fmt"

func shiftRight(array []int) []int {
	last := array[len(array)-1]

	for i := len(array) - 1; i > 0; i-- {
		array[i] = array[i-1]
	}

	array[0] = last

	return array
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	for _, number := range shiftRight(numbers) {
		fmt.Print(number, " ")
	}
	
	fmt.Print("\n")
}
