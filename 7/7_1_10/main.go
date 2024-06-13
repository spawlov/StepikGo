package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, k int

	_, _ = fmt.Scan(&n, &k)

	fmt.Println(combinations(n, k).String())
}

func combinations(n, k int) *big.Int {
	if n < k {
		return big.NewInt(0)
	}

	numerator := factorial(n)
	denomenator := new(big.Int).Mul(factorial(k), factorial(n-k))

	return new(big.Int).Div(numerator, denomenator)
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
