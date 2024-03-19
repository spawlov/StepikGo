package main

import . "fmt"

func main() {
	var n int
	_, _ = Scan(&n)
	factors := []int{}
    for n%2 == 0 {
		factors = append(factors, 2)
        n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
            factors = append(factors, i)
            n = n / i
        }
    }
	if n > 2 {
        factors = append(factors, n)
    }
	for _, factor := range factors {
		Print(factor, " ")
	}
	Println()
}