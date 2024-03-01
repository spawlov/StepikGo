package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if number > 0 {
		fmt.Println(1)
	} else if number < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}
}
