package main

import "fmt"

func main() {
	var one, two, three int

	_, _ = fmt.Scan(&one, &two, &three)

	fmt.Println(sumRange(one, two) + sumRange(two, three))
}

func sumRange(x, y int) int {
	if x > y {
		return 0
	}

	result := 0
	for i := x; i <= y; i++ {
		result += i
	}

	return result
}
