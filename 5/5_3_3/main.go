package main

import "fmt"

func findUniq(array []int) []int {
	counters := make(map[int]int)

	for _, number := range array {
		counters[number]++
	}

	uniqElems := []int{}
	for _, number := range array {
		if counters[number] == 1 {
			uniqElems = append(uniqElems, number)
		}
	}
	return uniqElems
}

func main() {
	var n int

	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	for _, number := range findUniq(numbers) {
		fmt.Print(number, " ")
	}
	fmt.Print("\n")
}
