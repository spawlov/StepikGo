package main

import "fmt"

func main() {
	var numOne, numTwo, numThree int

	_, _ = fmt.Scan(&numOne, &numTwo, &numThree)

	fmt.Println(doubleFactorial(numOne), doubleFactorial(numTwo), doubleFactorial(numThree))
}

func doubleFactorial(number int) int {
	var result int = 1

	for number > 0 {
		result *= number
		number -= 2
	}

	return result
}
