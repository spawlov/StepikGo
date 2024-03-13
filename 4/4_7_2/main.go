package main

import (
	. "fmt"
	. "math"
)

func countDivisors(num int) int {
	count := 0
	sqrt := int(Sqrt(float64(num)))
	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			count++
			if num/i != i {
				count++
			}
		}
	}
	return count
}

func main() {
	var a, b, k int
	_, _ = Scan(&a, &b, &k)
	for i := a; i <= b; i++ {
		if countDivisors(i) >= k {
			Printf("%d ", i)
		}
	}
}
