package main

import "fmt"

func main() {
	var (
		text string
		k    int
	)

	_, _ = fmt.Scan(&text)
	runes := []rune(text)

	_, _ = fmt.Scan(&k)

	result := "NO"
	for index, symbol := range runes {
		if index+1 == k {
			result = string(symbol)
			break
		}
	}
	fmt.Println(result)
}
