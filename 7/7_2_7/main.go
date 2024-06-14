package main

import "fmt"

func main() {
	var numOne, numTwo string

	_, _ = fmt.Scan(&numOne, &numTwo)
	switch {
	case len(numOne) > len(numTwo):
		fmt.Println(1)
	case len(numOne) < len(numTwo):
		fmt.Println(2)
	default:
		fmt.Println(0)
	}
}
