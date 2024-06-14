package main

import (
	"fmt"
	"math"
)

func main() {
	var number int

	_, _ = fmt.Scan(&number)

	fmt.Println(typeNumber(number))
}

func typeNumber(number int) string {
	if number < 2 {
		return "prime"
	}

	sqrt := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return "composite"
		}
	}
	return "prime"
}
