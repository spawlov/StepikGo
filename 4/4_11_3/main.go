package main

import (
	"fmt"
	"strconv"
)

func isPalindome(n int) string {
	if n < 10 {
		return "YES"
	} else {
		strN := strconv.Itoa(n)
		left := 0
		right := len(strN) - 1
		for left <= right {
			if strN[left] != strN[right] {
				return "NO"
			}
			left++
			right--
		}
		return "YES"
	}
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println(isPalindome(n))
}
