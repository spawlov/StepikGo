package main

import (
	. "fmt"
	. "math"
)

func main() {
	var n, digit float64
	_, _ = Scan(&n)
	result := Inf(-1)
	for i := 0.0; i < n; i++ {
		_, _ = Scan(&digit)
		if result < digit {
			result = digit
		}
	}
	Println(result)
}