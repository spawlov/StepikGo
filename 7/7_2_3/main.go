package main

import "fmt"

func main() {
	var n, m int

	_, _ = fmt.Scan(&n, &m)

	fmt.Println(numsAverage(n) + numsAverage(m))
}

func numsAverage(num int) float32 {
	var sum int

	for i := 1; i <= num; i++ {
		sum += i
	}
	return float32(sum) / float32(num)
}
