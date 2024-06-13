package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		number string
		result int = -1
	)
	_, _ = fmt.Scan(&number)

	ticketOne := isLuckyTicket(number)

	_, _ = fmt.Scan(&number)

	ticketTwo := isLuckyTicket(number)

	if ticketOne && ticketTwo {
		result = 1
	}

	fmt.Println(result)
}

func isLuckyTicket(number string) bool {
	hi := sumNumbers(convertToSlice(number)[0:3])
	low := sumNumbers(convertToSlice(number)[3:6])
	return hi == low
}

func sumNumbers(slice []rune) int {
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
