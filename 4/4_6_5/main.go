package main

import . "fmt"

func main() {
	var num, numbers, counter float64
	_, _ = Scan(&num)
	numbers, counter = 0, 0
	for num != 0 {
		numbers += num
		counter++
		Scan(&num)
	}
	Println(numbers / counter)
}
