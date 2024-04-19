package main

import (
	. "fmt"
	. "math"
)

func countSolutions(a, b, c, d, e float64) int {
	count := 0
	for x := 0.0; x <= 1000; x++ {
		if (a*Pow(x, 3)+b*Pow(x, 2)+c*x+d)/(x-e) == 0 {
			count++
		}
	}

	return count
}

func main() {
	var a, b, c, d, e float64
	Scan(&a, &b, &c, &d, &e)
	Println(countSolutions(a, b, c, d, e))
}
