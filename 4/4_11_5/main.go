package main

import "fmt"

func Pow(x, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

func main() {
	var x, n int
	_, _ = fmt.Scan(&x, &n)
	fmt.Println(Pow(x, n))
}
