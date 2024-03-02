package main

import "fmt"

func sumDigits(num int) int {
	if num < 10 {
		return num
	}
	return num%10 + sumDigits(num/10)
}

func main() {
	var x int
	fmt.Scan(&x)

	if sumDigits(x/1000) == sumDigits(x%1000) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
