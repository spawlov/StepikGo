package main

import "fmt"

func main() {
	var n, number, counter int

	_, _ = fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&number)
		if number == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
