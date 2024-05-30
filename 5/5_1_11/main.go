package main

import "fmt"

func swapNeighborElements(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func main() {
	var N int

	_, _ = fmt.Scan(&N)

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	swapNeighborElements(numbers)

	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
}
