package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numOne, numTwo string

	_, _ = fmt.Scan(&numOne, &numTwo)
	switch {
	case sumDigits(convertToSlice(numOne)) > sumDigits(convertToSlice(numTwo)):
		fmt.Println(1)
	case sumDigits(convertToSlice(numOne)) < sumDigits(convertToSlice(numTwo)):
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}

func sumDigits(slice []rune) int {
	sum := 0
	for _, char := range slice {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	return sum
}

func convertToSlice(text string) []rune {
	result := []rune{}
	for _, char := range text {
		result = append(result, char)
	}
	return result
}
