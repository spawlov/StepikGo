package main

import . "fmt"

func main() {
	var a, b string
	_, _ = Scan(&a)
	for i := 1; 0 < i; i++ {
		_, _ = Scan(&b)
		if a == b {
			Println(i)
			break
		}
	}
}
