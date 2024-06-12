package main

import "fmt"

func main() {
	var (
		n, count int
		filename string
	)

	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&filename)

	for i := 2; i < n; i++ {
		if filename[i-2] == 'x' && filename[i-1] == 'x' && filename[i] == 'x' {
			count++
		}
	}
	fmt.Println(count)
}
