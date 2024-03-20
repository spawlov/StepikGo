package main

import (
	. "fmt"
	. "math"
)

func solveEquation(a, b, c, d float64) []int {
	result := []int{}
	for x := 0.0; x <= 1000; x++ {
		if a*Pow(x, 3)+b*Pow(x, 2)+c*x+d == 0 {
			result = append(result, int(x))
		}
	}
	return result
}

func main() {
	var a, b, c, d float64
	_, _ = Scan(&a, &b, &c, &d)
	results := solveEquation(a, b, c, d)
	if len(results) > 0 {
		for _, result := range results {
			Print(result, " ")
		}
		Println()
	}
}
