package main

import "fmt"

func main() {
	var numOne, numTwo int

	_, _ = fmt.Scan(&numOne, &numTwo)

	fmt.Println(reverseNumber(numOne) + reverseNumber(numTwo))
}

func reverseNumber(number int) int {
	var result int

	if number < 10 {
		return number
	}

	for number > 0 {
		remainder := number % 10
		result = result*10 + remainder
		number /= 10
	}

	return result
}
