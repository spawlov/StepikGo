package main

import "fmt"

func main() {
	var n, h, width int

	_, _ = fmt.Scan(&n, &h)

	heights := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&heights[i])
		if heights[i] > h {
			width++
		}
		width++
	}
	
	fmt.Println(width)
}
