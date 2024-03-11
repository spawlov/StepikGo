package main

import . "fmt"

func main() {
	var num int
	_, _ = Scan(&num)
	positive, negative := 0, 0
	for num != 0 {
		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		}
		Scan(&num)
	}
	Println(positive - negative)
}
