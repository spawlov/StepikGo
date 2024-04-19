package main

import (
	. "fmt"
	. "math"
)

func main() {
	var N int
	Scan(&N)

	maxNum := MinInt64
	minNum := MaxInt64

	for i := 0; i < N; i++ {
		var num int
		Scan(&num)
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	Println(maxNum - minNum)
}
