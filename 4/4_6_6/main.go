package main

import . "fmt"

func main() {
	var num, previous, counter int
	_, _ = Scan(&num)
	for num != 0 {
		if num > previous {
			counter++
		}
		previous = num
		_, _ = Scan(&num)
	}
	Println(counter - 1)
}
