package main

import (
	"fmt"
	"strings"
)

func main() {
	var counter, n int

	_, _ = fmt.Scan(&n)

	var numberId []int
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
		if numbers[i]%3 == 0 && numbers[i]%10 == 7 {
			counter++
			numberId = append(numberId, i)
		}
	}

	for _, id := range numberId {
		numbers[id] = counter
	}

	fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))
}
