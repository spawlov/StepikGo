package main

import "fmt"

func main() {
	var (
		firstChar, lastChar string
		start, end          rune
	)

	_, _ = fmt.Scan(&firstChar)
	_, _ = fmt.Scan(&lastChar)

	if firstChar[0] > lastChar[0] {
		start = rune(lastChar[0])
		end = rune(firstChar[0])
	} else {
		end = rune(lastChar[0])
		start = rune(firstChar[0])
	}

	for i := start; i <= end; i++ {
		fmt.Print(string(i), " ")
	}
}
