package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	strN := strconv.Itoa(n)
	resultStr := ""
	for _, digit := range strN {
		if digit != '5' && digit != '7' {
			resultStr += string(digit)
		}
	}
	result, _ := strconv.Atoi(resultStr)
	fmt.Println(result)
}
