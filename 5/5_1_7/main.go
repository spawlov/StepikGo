package main

import "fmt"

func main() {
	var size int

	_, _ = fmt.Scan(&size)

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	for i, value := range numbers {
		if i%3 == 0 {
			fmt.Print(value, " ")
		}
	}
	fmt.Print("\n")
}
