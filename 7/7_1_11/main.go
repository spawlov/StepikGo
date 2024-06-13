package main

import "fmt"

func main() {
	var a, b int

	_, _ = fmt.Scan(&a, &b)

	fmt.Println(getSign(a) + getSign(b))
}

func getSign(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}
