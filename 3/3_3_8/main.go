package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	d := math.Pow(b, 2) - 4*a*c
	if d == 0 {
		fmt.Println(-b / (a * 2))
	} else if d > 0 {
		x1 := (-b - math.Sqrt(d)) / (a * 2)
		x2 := (-b + math.Sqrt(d)) / (a * 2)
		if x1 > x2 {
			fmt.Println(x2, "\n", x1)
		} else {
			fmt.Println(x1, "\n", x2)
		}
	}
}
