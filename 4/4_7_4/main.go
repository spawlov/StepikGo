package main

import . "fmt"

func sumDigits(num int) int {
	result := 0
	for num > 0 {
		result += num % 10
		num /= 10
	}
	return result
}

func main() {
	var number int
	_, _ = Scan(&number)
	maxSum, maxNum := 0, 0
	for i := 0; i <= number; i++ {
		currSum := sumDigits(i)
		if currSum > maxSum {
			maxSum = currSum
			maxNum = i
		} else if currSum == maxSum && i < maxNum {
			maxNum = i
		}
	}
	Printf("%d %d \n", maxNum, maxSum)
}
