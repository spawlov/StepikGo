package main

import "fmt"

func digitalRoot(n int) int {
	var sum int

	if n < 10 {
		return n
	}

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return digitalRoot(sum)
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	result := digitalRoot(n)
	fmt.Println(result)
}
