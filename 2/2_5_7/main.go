package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Scan(&number)

	result := number % 10
	result = result + (number/10)%10
	result = result + number/100

	fmt.Println(result)
}
