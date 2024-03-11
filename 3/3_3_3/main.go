package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if (number >= -3 && number <= 1) || (number >= 5 && number <= 9) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
