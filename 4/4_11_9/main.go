package main

import (
	"fmt"
)

func hasDistinctDigits(num int) bool {
	digits := [10]bool{}
	for num > 0 {
		digit := num % 10
		if digits[digit] {
			return false
		}
		digits[digit] = true
		num /= 10
	}
	return true
}

func main() {
	var number int

	_, _ = fmt.Scan(&number)

	for {
		number++
		if hasDistinctDigits(number){
			fmt.Println(number)
			break
		}
	}
}
