package main

import . "fmt"

func main() {
	var number, previous, changes int
	_, _ = Scan(&number)
	for number != 0 {
		if (number > 0 && previous < 0) || (number < 0 && previous > 0) {
			changes++
		}
		previous = number
		_, _ = Scan(&number)
	}
	Println(changes)
}
