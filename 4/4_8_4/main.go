package main

import . "fmt"

func main() {
	var n int
	for n <= 100 {
		_, _ = Scan(&n)
		if n < 10 || n > 100 {
			continue
		}
		Println(n)
	}
}
